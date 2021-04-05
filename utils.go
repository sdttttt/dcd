package huck

import (
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
