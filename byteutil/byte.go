package byteutil

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Split 数据分片
func Split(buf []byte, limit int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/limit+1)
	for len(buf) >= limit {
		chunk, buf = buf[:limit], buf[limit:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

// Join 数据合并
func Join(s [][]byte) []byte {
	return bytes.Join(s, []byte(""))
}

// JpegToBytes ...
func JpegToBytes(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)
	return buf.Bytes(), err
}

// BytesToJpeg ...
func BytesToJpeg(buf []byte) (image.Image, error) {
	return jpeg.Decode(bytes.NewReader(buf))
}

// PngToBytes ...
func PngToBytes(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	return buf.Bytes(), err
}

// BytesToPng ...
func BytesToPng(buf []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(buf))
}

// IOReaderToBytes read contents from io.Reader
func IOReaderToBytes(r io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// BytesToIOReader bytes => io.Reader
func BytesToIOReader(buf []byte) io.Reader {
	return bytes.NewReader(buf)
}
