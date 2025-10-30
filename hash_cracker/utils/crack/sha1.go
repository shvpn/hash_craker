package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashSHA1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
