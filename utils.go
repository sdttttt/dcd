package huck

import (
	"os"
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytes := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bytes))
}

func BytesToString(b []byte) string {
	bytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	str := reflect.StringHeader{
		Data: bytes.Data,
		Len:  bytes.Len,
	}
	return *(*string)(unsafe.Pointer(&str))
}

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}
