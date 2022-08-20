package byteutils

import (
	"bytes"
	"testing"
)

func TestStringToByte(t *testing.T) {
	x := "Hello Gopher!"
	y := StringToByte(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
		t.Fail()
	}
}

func TestByteToString(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := ByteToString(x)
	z := string(x)

	if y != z {
		t.Fail()
	}
}

func TestUint64ToByte(t *testing.T) {
	x := uint64(1234567890)
	y := Uint64ToByte(x)
	t.Logf("Uint64ToBytes: %b", y)
}

func TestByteToUint64(t *testing.T) {
	x := uint64(1234567890)
	y := Uint64ToByte(x)
	z := ByteToUint64(y)
	t.Logf("%d", z)
}
