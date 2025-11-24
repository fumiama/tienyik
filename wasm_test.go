package tienyik

import (
	"context"
	"testing"
)

func TestSigner(t *testing.T) {
	ctx := context.Background()
	sg := NewSigner(ctx)
	sigstr := sg.GenKeyNew(
		ctx, 60, 1763891935806, 1763891937568, "9c047f01dfab388a3ef7c4ae34855e3a",
		"6806caebc2cdaef5f10987a94d21cf1f", "/api/cdserv/client/msgcenter/page",
		"8aff800b3c96e620043bb7deb4a0258d", 103010001,
	)
	if sigstr != "66443F2BF714E1D2B4AC8D3B3CBC3913" {
		t.Fatal("got", sigstr)
	}
	t.Fail()
}
