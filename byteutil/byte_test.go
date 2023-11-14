package byteutil

import (
	"bytes"
	"encoding/base64"
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

// go test -bench="." -benchmem
// 标准转换string()
func Benchmark_NormalBytesToString(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

// 强转换[]byte到string
func Benchmark_BytesToString(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = BytesToString(x)
	}
}

// 标准转换[]byte
func Benchmark_NormalStringToBytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

// 强转换string到[]byte
func Benchmark_StringToBytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(x)
	}
}

func TestPngToBytes(t *testing.T) {
	s := "iVBORw0KGgoAAAANSUhEUgAAAMgAAADIAQMAAACXljzdAAAABlBMVEX///8AAABVwtN+AAAA9UlEQVR4nOyWMY6EMBRDjShS5gg5CkeDo+UoHCElBcIrZ4KWFen3/1FczCA9GpP/Y2No6EsVKZVEXgj1+fBAdgChYNnqHzD7IAuvUFJeeYay8PJFcMEdwbL5InV2EivpTNV/kraNeZ3O/p5aJE0pr+jLJpETkvfv74RYJnHXRMcdOoUCTPRAAKwkd3ycPJ0aJtpTHnLFU29NdEGYgSA/umSUkHBAmhLJE5Hbn/MxSx7taT6ich0eSMufrOSJfJ6PAXLndk2Z7dXsLJPakaBvTUek+oGaUtepPXK3J8wswHuqTJLPNgLQHaIO5YEMDfnUTwAAAP//nDsYE5ec1W8AAAAASUVORK5CYII="
	b, _ := base64.StdEncoding.DecodeString(s)
	// 转Png Succeed
	i, err := BytesToPng(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Result: BytesToPng Succeed")
	// 转Jpeg Failed
	ii, err := BytesToJpeg(b)
	if err != nil {
		t.Log(err)
	}
	if ii == nil {
		t.Log("Result: BytesToJpeg Failed")
	}
	// Png转Bytes
	n, err := PngToBytes(i)
	if err != nil {
		t.Fatal(err)
	}
	if s == base64.StdEncoding.EncodeToString(n) {
		t.Log("Result: PngToBytes Succeed")
	}
	// Jpg转Bytes
	nn, err := JpegToBytes(i)
	if err != nil {
		t.Fatal(err)
	}
	if nn != nil {
		t.Log("Result: JpegToBytes Succeed")
	}
}

func TestIOReaderToBytes(t *testing.T) {
	b := []byte("hello")
	r := BytesToIOReader(b)
	bb, err := IOReaderToBytes(r)
	t.Log(string(bb), err)
}
