package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// Returns the lower-case hex string (e.g., "6a85df...")
func Hash(text string) string {
	sum := md5.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}
