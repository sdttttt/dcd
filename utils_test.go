package huck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBytes(t *testing.T) {
	str := "abcdefghijklmnopqrstuvwxy"

	byte1 := []byte(str)
	byte2 := StringToBytes(str)

	assert.Equal(t, byte1, byte2, "StringToBytes should ok.")
}
