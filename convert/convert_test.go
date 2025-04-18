package convert

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestAnyToString(t *testing.T) {
	assert.Equalf(t, "0", AnyToString(0), "FAIL")
	assert.Equalf(t, "-123", AnyToString(-123), "FAIL")
	assert.Equalf(t, "-123.123456", AnyToString(-123.123456), "FAIL")
	assert.Equalf(t, "9223372036854775807", AnyToString(math.MaxInt64), "FAIL")
}

func TestIntToString(t *testing.T) {
	assert.Equalf(t, "12345678", IntToString(12345678), "FAIL")
	assert.Equalf(t, "-12345678", IntToString(-12345678), "FAIL")
}

func TestInt64ToString(t *testing.T) {
	assert.Equalf(t, "12345678", Int64ToString(int64(12345678)), "FAIL")
}

func TestUint64ToString(t *testing.T) {
	assert.Equalf(t, "12345678", Uint64ToString(uint64(12345678)), "FAIL")
}

func TestFloat64ToString(t *testing.T) {
	assert.Equalf(t, "1.23456789", Float64ToString(1.23456789), "FAIL")
}

func TestFloat32ToString(t *testing.T) {
	assert.Equalf(t, "1.2345679", Float32ToString(float32(1.23456789)), "FAIL")
}

func TestStringToFloat64(t *testing.T) {
	assert.Equalf(t, 1.23456789, StringToFloat64("1.23456789"), "FAIL")
}

func TestStringToFloat32(t *testing.T) {
	assert.Equalf(t, float32(1.2345679), StringToFloat32("1.23456789"), "FAIL")
}

func TestStringToInt(t *testing.T) {
	assert.Equalf(t, 1234567, StringToInt("1234567"), "FAIL")
}

func TestStringToUint(t *testing.T) {
	assert.Equalf(t, uint(1234567), StringToUint("1234567"), "FAIL")
}

func TestStringToInt32(t *testing.T) {
	assert.Equalf(t, int32(1234567), StringToInt32("1234567"), "FAIL")
}

func TestStringToUint32(t *testing.T) {
	assert.Equalf(t, uint32(math.MaxUint32), StringToUint32("123456789011"), "FAIL")
}

func TestStringToInt64(t *testing.T) {
	assert.Equalf(t, int64(1234567890123), StringToInt64("1234567890123"), "FAIL")
}

func TestStringToUint64(t *testing.T) {
	assert.Equalf(t, uint64(1234567890123), StringToUint64("1234567890123"), "FAIL")
}

func TestIntToUint(t *testing.T) {
	assert.Equalf(t, uint(1234567), IntToUint(1234567), "FAIL")
}

func TestUintToInt(t *testing.T) {
	assert.Equalf(t, 1234567, UintToInt(1234567), "FAIL")
}

func TestJsonNumberToInt(t *testing.T) {
	type Data struct {
		Number json.Number `json:"number"`
	}
	var d Data
	_ = json.Unmarshal([]byte(`{"number":"1234567""}`), &d)
	_ = json.Unmarshal([]byte(`{"number":1234567}`), &d)
	assert.Equalf(t, 1234567, JsonNumberToInt(d.Number), "FAIL")
}

func TestMapToJson(t *testing.T) {
	j, _ := MapToJson(map[string]string{"a": "b", "c": "d"})
	assert.Equalf(t, "{\"a\":\"b\",\"c\":\"d\"}", j, "FAIL")
}

func TestJsonToMap(t *testing.T) {
	m, _ := JsonToMap("{\"a\":\"b\",\"c\":\"d\"}")
	assert.Equalf(t, "b", m["a"], "FAIL")
	assert.Equalf(t, "d", m["c"], "FAIL")
}

func TestBase64Encode(t *testing.T) {
	assert.Equalf(t, "aGVsbG8=", Base64Encode([]byte("hello")), "FAIL")
}

func TestBase64Decode(t *testing.T) {
	b, _ := Base64Decode("aGVsbG8=")
	assert.Equalf(t, "hello", string(b), "FAIL")
}

func TestBase64UrlEncode(t *testing.T) {
	s := "钟山风雨起苍黄百万雄师过大江"
	r := Base64UrlEncode([]byte(s))
	_, err := Base64Decode(r)
	assert.Equalf(t, "6ZKf5bGx6aOO6Zuo6LW36IuN6buE55m-5LiH6ZuE5biI6L-H5aSn5rGf", r, "FAIL")
	assert.NotEqualf(t, nil, err, "FAIL")
	rr, err := Base64UrlDecode(r)
	assert.Equalf(t, s, string(rr), "FAIL")
}

func TestBoolToInt(t *testing.T) {
	assert.Equalf(t, 1, BoolToInt(true), "FAIL")
	assert.Equalf(t, 0, BoolToInt(false), "FAIL")
}

func TestIntToBool(t *testing.T) {
	assert.Equalf(t, true, IntToBool(1), "FAIL")
	assert.Equalf(t, false, IntToBool(0), "FAIL")
}

func TestToJSON(t *testing.T) {
	msg := struct {
		Id      int32
		Content string
	}{
		Id:      1,
		Content: "hello",
	}
	assert.Equalf(t, "{\"Id\":1,\"Content\":\"hello\"}", ToJSON(msg), "FAIL")
}

func TestToJSONBytes(t *testing.T) {
	msg := struct {
		Id      int32
		Content string
	}{
		Id:      1,
		Content: "hello",
	}
	assert.Equalf(t, "{\"Id\":1,\"Content\":\"hello\"}", string(ToJSONBytes(msg)), "FAIL")
}

func TestToJSONRaw(t *testing.T) {
	msg := struct {
		Id      int32
		Content string
	}{
		Id:      1,
		Content: "hello",
	}
	assert.Equalf(t, "{\"Id\":1,\"Content\":\"hello\"}", string(ToJSONRaw(msg)), "FAIL")
}

func TestCopyProperties(t *testing.T) {
	msg := struct {
		Id      int32
		Content string
	}{
		Id:      1,
		Content: "hello",
	}

	type Msg struct {
		Id      int32
		Content string
		Title   string
	}
	var x = struct {
		Id      int32
		Content string
		Title   string
	}{
		Title: "title",
	}
	err := CopyProperties(msg, &x)
	assert.Equalf(t, nil, err, "FAIL")
	assert.Equalf(t, "hello", x.Content, "FAIL")
}
