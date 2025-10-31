package hcli

import (
	"strings"

	base14 "github.com/fumiama/go-base16384"
	"github.com/fumiama/tienyik/internal/log"
)

var eps = base14.DecodeString("栝啇俌蠯姙呗宬籣欞敖蚹煮岎冃勀紀㴆")

func ep(p string) string {
	sb := &strings.Builder{}
	sb.WriteString(eps)
	sb.WriteString(p)
	log.Debugln("ep wraps:", sb)
	return sb.String()
}
