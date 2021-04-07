package huck

import (
	"os"
	"reflect"
	"unsafe"
)

// StringToBytes to string to []byte.
func StringToBytes(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytes := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bytes))
}

// BytesToString to []byte to string.
func BytesToString(b []byte) string {
	bytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	str := reflect.StringHeader{
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
