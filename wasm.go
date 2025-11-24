package tienyik

import (
	"context"
	"fmt"
	"strings"

	_ "embed"

	"github.com/fumiama/tienyik/internal/op"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

//go:embed main.1755740488270.wasm
var wasmdata []byte

type Signer struct {
	rt               wazero.Runtime
	md               api.Module
	genKey           api.Function
	genKeyNew        api.Function
	genKeyWithoutURI api.Function
	_malloc          api.Function
	_free            api.Function
}

func NewSigner(ctx context.Context) (sg Signer) {
	rt := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())
	_, err := rt.NewHostModuleBuilder("a").
		// ___cxa_throw
		NewFunctionBuilder().WithFunc(func(ctx context.Context, e, n, t uint32) {
		panic(fmt.Sprintf("___cxa_throw: ptr(e)=%08x, type(n)=%08x, destructor(t)=%08x", e, n, t))
	}).Export("c").
		// __abort_js
		NewFunctionBuilder().WithFunc(func(ctx context.Context) {
		panic("wasm aborted")
	}).Export("a").
		// _emscripten_resize_heap
		NewFunctionBuilder().WithFunc(func(ctx context.Context, _ uint32) uint32 {
		panic("wasm oom")
	}).Export("b").
		Instantiate(ctx)
	if err != nil {
		panic(err)
	}

	md, err := rt.InstantiateWithConfig(ctx, wasmdata,
		wazero.NewModuleConfig())
	if err != nil {
		_ = rt.Close(ctx)
		panic(err)
	}

	sg.rt = rt
	sg.md = md

	sg.genKey = md.ExportedFunction("f")
	sg.genKeyNew = md.ExportedFunction("g")
	sg.genKeyWithoutURI = md.ExportedFunction("h")
	sg._malloc = md.ExportedFunction("i")
	sg._free = md.ExportedFunction("j")

	return
}

// GenKey is the go repr of js func
//
//	generatorSign(e) {
//		const t = Module.lengthBytesUTF8(e.secretKey) + 1
//			, n = Module._malloc(t);
//		Module.stringToUTF8(e.secretKey, n, t);
//		const r = Module._gen_key(e.deviceType, BigInt(e.timestamp), BigInt(e.requestId), n, e.tenantId, e.userId, e.version)
//			, o = Module.UTF8ToString(r);
//		return Module._free(n),
//		Module._free(r),
//		o
//	}
func (sg *Signer) GenKey(
	ctx context.Context, deviceType, timestamp, requestID uint64,
	secretKey string, tenantID, userID, version uint64,
) string {
	t := len(secretKey) + 1
	n := sg.malloc(ctx, uint64(t))
	if !sg.md.Memory().WriteString(uint32(n), secretKey+"\x00") {
		panic("write out-of-bound")
	}
	defer sg.free(ctx, n)

	return sg.string(op.Must(sg.genKey.Call(
		ctx, deviceType, timestamp, requestID,
		n, tenantID, userID, version),
	)[0])
}

// GenKeyNew is the go repr of js func
//
//	generatorSignNew(e) {
//		const t = Module.lengthBytesUTF8(e.secretKey) + 1
//			, n = Module._malloc(t);
//		Module.stringToUTF8(e.secretKey, n, t);
//		const r = Module.lengthBytesUTF8(e.userEid) + 1
//			, o = Module._malloc(r);
//		Module.stringToUTF8(e.userEid, o, r);
//		const i = Module.lengthBytesUTF8(e.requestUri) + 1
//			, a = Module._malloc(i);
//		Module.stringToUTF8(e.requestUri, a, i);
//		const s = Module.lengthBytesUTF8(this.serverNode) + 1
//			, l = Module._malloc(s);
//		Module.stringToUTF8(this.serverNode, l, s);
//		const c = Module._gen_key_new(e.deviceType, BigInt(e.timestamp), BigInt(e.requestId), n, o, a, l, e.version)
//			, u = Module.UTF8ToString(c);
//		return Module._free(n),
//		Module._free(o),
//		Module._free(a),
//		Module._free(l),
//		Module._free(c),
//		u
//	}
func (sg *Signer) GenKeyNew(
	ctx context.Context, deviceType, timestamp, requestID uint64,
	secretKey, userEID, requestURI, serverNode string, version uint64,
) string {
	t := len(secretKey) + 1
	n := sg.malloc(ctx, uint64(t))
	if !sg.md.Memory().WriteString(uint32(n), secretKey+"\x00") {
		panic("write out-of-bound")
	}
	defer sg.free(ctx, n)

	r := len(userEID) + 1
	o := sg.malloc(ctx, uint64(r))
	if !sg.md.Memory().WriteString(uint32(o), userEID+"\x00") {
		panic("write out-of-bound")
	}
	defer sg.free(ctx, o)

	i := len(requestURI) + 1
	a := sg.malloc(ctx, uint64(i))
	if !sg.md.Memory().WriteString(uint32(a), requestURI+"\x00") {
		panic("write out-of-bound")
	}
	defer sg.free(ctx, a)

	s := len(serverNode) + 1
	l := sg.malloc(ctx, uint64(s))
	if !sg.md.Memory().WriteString(uint32(l), serverNode+"\x00") {
		panic("write out-of-bound")
	}
	defer sg.free(ctx, l)

	return sg.string(op.Must(sg.genKeyNew.Call(
		ctx, deviceType, timestamp, requestID,
		n, o, a, l, version),
	)[0])
}

func (sg *Signer) malloc(ctx context.Context, n uint64) uint64 {
	return op.Must(sg._malloc.Call(ctx, n))[0]
}

func (sg *Signer) free(ctx context.Context, n uint64) {
	op.Must(sg._free.Call(ctx, n))
}

func (sg *Signer) string(ptr uint64) string {
	buf := strings.Builder{}
	x := uint32(ptr)
	for {
		b, ok := sg.md.Memory().ReadByte(x)
		x++
		if !ok {
			panic("read out-of-bound")
		}
		if b == 0 {
			break
		}
		buf.WriteByte(b)
	}
	return buf.String()
}

func (sg *Signer) IsClosed() bool {
	return sg.md == nil || sg.rt == nil || sg.md.IsClosed() ||
		sg.genKey == nil || sg.genKeyNew == nil || sg.genKeyWithoutURI == nil ||
		sg._malloc == nil || sg._free == nil
}

func (sg *Signer) Close(ctx context.Context) {
	if sg.md != nil && !sg.md.IsClosed() {
		sg.md.Close(ctx)
		sg.md = nil
	}
	if sg.rt != nil {
		sg.rt.Close(ctx)
		sg.rt = nil
	}
}
