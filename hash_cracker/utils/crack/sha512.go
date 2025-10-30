package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashSHA512(text string) string {
	h := sha512.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
