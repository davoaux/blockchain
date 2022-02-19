package util

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Encode input has a hexadecimal string
func Encode(input string) string {
	builder := strings.Builder{}
	builder.Grow(64)
	checksum := sha256.Sum256([]byte(input))
	for _, n := range checksum {
		builder.WriteString(fmt.Sprintf("%02x", n))
	}
	return builder.String()
}
