package byteutils

import (
	"bytes"
	"testing"
)

func TestString2Bytes(t *testing.T) {
	x := "Hello Gopher!"
	y := String2Bytes(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
		t.Fail()
	}
}

func TestBytes2String(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := Bytes2String(x)
	z := string(x)

	if y != z {
		t.Fail()
	}
}

func TestUint64ToBytes(t *testing.T) {
	x := uint64(1234567890)
	y := Uint64ToBytes(x)
	t.Logf("Uint64ToBytes: %b", y)
}

func TestBytesToUint64(t *testing.T) {
	x := uint64(1234567890)
	y := Uint64ToBytes(x)
	z := BytesToUint64(y)
	t.Logf("%d", z)
}
