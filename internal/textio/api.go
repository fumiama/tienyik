package textio

import (
	"runtime"
	"strings"
)

func API() string {
	pc, f, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get api of caller")
	}
	if strings.Contains(f, "\\") {
		f = strings.ReplaceAll(f, "\\", "/")
	}
	_, p, ok := strings.Cut(f, "/tienyik/")
	if !ok {
		panic("cannot cut api " + f + " of caller")
	}
	f = strings.TrimSuffix(p, ".go")
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		panic("cannot get func name of caller, api: " + f)
	}
	p = fn.Name()
	i := strings.LastIndex(p, ".")
	if i < 0 {
		panic("func name of caller '" + p + " has no '.', api: " + f)
	}
	p = p[i+1:]
	if len(p) <= 1 {
		panic("func name of caller '" + p + " too short', api: " + f)
	}
	sb := strings.Builder{}
	sb.WriteString(f)
	sb.WriteByte('/')
	sb.WriteString(strings.ToLower(p[:1]))
	sb.WriteString(p[1:])
	return sb.String()
}
