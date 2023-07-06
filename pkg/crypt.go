package goshc

import (
	"crypto/rand"
	"crypto/rc4"
	"fmt"
)

func Crypt(data, key []byte) error {

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return err
	}

	cipher.XORKeyStream(data, data)

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
