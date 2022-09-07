package byteutil

import (
	"bytes"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	x := "Hello Gopher!"
	y := StringToBytes(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
		t.Fail()
	}
}

func TestBytesToString(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := BytesToString(x)
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

func TestSplit(t *testing.T) {
	slice := Split([]byte("Hello Gopher!"), 1)
	t.Logf("%s", slice)
}

func TestJoin(t *testing.T) {
	b := Join(Split([]byte("Hello Gopher!"), 1))
	t.Logf("%s", b)
}
