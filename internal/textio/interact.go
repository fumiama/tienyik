package textio

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func NoEchoScanln(a ...any) (n int, err error) {
	fd := int(os.Stdin.Fd())
	bytePassword, err := term.ReadPassword(fd)
	if err != nil {
		return 0, err
	}
	return fmt.Sscanln(string(bytePassword)+"\n", a...)
}
