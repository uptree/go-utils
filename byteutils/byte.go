package byteutils

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

// StringToByte 强制转换 []byte(s)
func StringToByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// ByteToString 强制转换 string(s)
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Uint64ToByte uint64转byte
func Uint64ToByte(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// ByteToUint64 byte转uint64
func ByteToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
