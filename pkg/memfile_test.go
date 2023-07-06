package goshc

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemfile(t *testing.T) {

	data := []byte("TestingData")

	file, err := MemOpen(data)
	assert.Nil(t, err)

	readData, err := io.ReadAll(file)
	assert.Nil(t, err)
	assert.Equal(t, data, readData)
}
