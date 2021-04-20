package stuff

import (
	"crypto/sha256"
	"encoding/hex"
)
func EncryptData(data string) string {
	sum := sha256.Sum256([]byte(data)) // we encript the data

	return hex.EncodeToString(sum[:])
}