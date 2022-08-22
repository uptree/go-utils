package byteutil

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

// StringToBytes 强制转换 []byte(s)
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// BytesToString 强制转换 string(s)
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Uint64ToBytes uint64转byte
func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// BytesToUint64 byte转uint64
func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
