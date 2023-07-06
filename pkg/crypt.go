package goshc

import (
	"crypto/rand"
	"fmt"
)

// Replace this with something later
// Simple xor across the key
func Crypt(data, key []byte) error {

	for index, char := range data {
		modPosition := index
		if index >= len(key) {
			modPosition = index % len(key)
		}
		data[index] = char ^ key[modPosition]
	}

	return nil
}

func ToGoString(data []byte) string {
	var out string
	for _, char := range data {
		if len(out) > 0 {
			out = fmt.Sprintf("%s, ", out)
		}
		out = fmt.Sprintf("%s0x%02x", out, char)
	}

	return out
}

func GenerateKey() []byte {
	token := make([]byte, 32)
	rand.Read(token)

	return token
}
