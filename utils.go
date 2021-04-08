package huck

import (
	"os"
	"unsafe"
)

// StringHeader Don't uintptr
type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// SliceHeader Don't uintptr
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// StringToBytes to string to []byte.
func StringToBytes(s string) []byte {
	str := (*StringHeader)(unsafe.Pointer(&s))
	bytes := SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bytes))
}

// BytesToString to []byte to string.
func BytesToString(b []byte) string {
	bytes := (*SliceHeader)(unsafe.Pointer(&b))

	str := StringHeader{
		Data: bytes.Data,
		Len:  bytes.Len,
	}
	return *(*string)(unsafe.Pointer(&str))
}

// IsFileExist to quick check file is exist.
func IsFileExist(filename string) bool {
	_, err := os.Lstat(filename)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
