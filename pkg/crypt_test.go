package goshc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypt(t *testing.T) {

	original := []byte("TestingData")
	test := make([]byte, len(original))
	copy(test, original)
	key := []byte("TestKey")

	err := Crypt(test, key)
	assert.Nil(t, err)
	assert.NotEqual(t, original, test)

	err = Crypt(test, key)
	assert.Nil(t, err)
	assert.Equal(t, original, test)
}

func TestToGoString(t *testing.T) {

	expected := "0x00, 0x01, 0x02"
	test := []byte{0x00, 0x01, 0x02}

	output := ToGoString(test)
	assert.Equal(t, expected, output)
}
