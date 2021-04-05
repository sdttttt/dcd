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

func TestBytesToString(t *testing.T) {
	str := "abcdefghijklmnopqrstuvwxy"

	bytes := []byte(str)

	str3 := string(bytes)
	str2 := BytesToString(bytes)

	assert.Equal(t, str3, str2, "StringToBytes should ok.")
}

func TestIsFileExists(t *testing.T) {
	filename := "README.md"
	assert.True(t, IsFileExist(filename), "README.md is exist.")
}
