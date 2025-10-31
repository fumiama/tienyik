package tienyik

import (
	"crypto/sha256"
	"encoding/hex"
)

func ChallengePassword(pwd string, chlg string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	h.Write([]byte(chlg))
	var sum [sha256.Size]byte
	return hex.EncodeToString(h.Sum(sum[:0]))
}

func ChallengeSHA256Password(pwd string, chlg string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	var sum [sha256.Size]byte
	s := hex.EncodeToString(h.Sum(sum[:0]))
	h.Reset()
	h.Write([]byte(s))
	h.Write([]byte(chlg))
	return hex.EncodeToString(h.Sum(sum[:0]))
}
