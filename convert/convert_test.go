package convert

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyToString(t *testing.T) {
	assert.Equalf(t, "0", AnyToString(0), "failed")
	assert.Equalf(t, "-123", AnyToString(-123), "failed")
	assert.Equalf(t, "-123.123456", AnyToString(-123.123456), "failed")
	assert.Equalf(t, "9223372036854775807", AnyToString(math.MaxInt64), "failed")
}

func TestIntToString(t *testing.T) {
	assert.Equalf(t, "12345678", IntToString(12345678), "failed")
	assert.Equalf(t, "-12345678", IntToString(-12345678), "failed")
}

func TestUint64ToString(t *testing.T) {
	assert.Equalf(t, "12345678", Uint64ToString(uint64(12345678)), "failed")
}

func TestFloat64ToString(t *testing.T) {
	assert.Equalf(t, "1.23456789", Float64ToString(1.23456789), "failed")
}

func TestFloat32ToString(t *testing.T) {
	assert.Equalf(t, "1.2345679", Float32ToString(float32(1.23456789)), "failed")
}

func TestStringToInt(t *testing.T) {
	assert.Equalf(t, 1234567, StringToInt("1234567"), "failed")
}

func TestStringToInt32(t *testing.T) {
	assert.Equalf(t, int32(1234567), StringToInt32("1234567"), "failed")
}

func TestStringToInt64(t *testing.T) {
	assert.Equalf(t, int64(1234567890123), StringToInt64("1234567890123"), "failed")
}

func TestStringToUint64(t *testing.T) {
	assert.Equalf(t, uint64(1234567890123), StringToUint64("1234567890123"), "failed")
}

func TestIntToUint(t *testing.T) {
	assert.Equalf(t, uint(1234567), IntToUint(1234567), "failed")
}

func TestUintToInt(t *testing.T) {
	assert.Equalf(t, 1234567, UintToInt(1234567), "failed")
}

func TestMapToJson(t *testing.T) {
	j, _ := MapToJson(map[string]string{"a": "b", "c": "d"})
	assert.Equalf(t, "{\"a\":\"b\",\"c\":\"d\"}", j, "failed")
}

func TestJsonToMap(t *testing.T) {
	m, _ := JsonToMap("{\"a\":\"b\",\"c\":\"d\"}")
	assert.Equalf(t, "b", m["a"], "failed")
	assert.Equalf(t, "d", m["c"], "failed")
}

func TestBase64Encode(t *testing.T) {
	assert.Equalf(t, "aGVsbG8=", Base64Encode([]byte("hello")), "failed")
}

func TestBase64Decode(t *testing.T) {
	b, _ := Base64Decode("aGVsbG8=")
	assert.Equalf(t, "hello", string(b), "failed")
}
