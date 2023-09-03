package convert

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// AnyToString ...
func AnyToString(i interface{}) string {
	var s string
	switch v := i.(type) {
	case nil:
		s = ""
	case int:
		s = strconv.Itoa(v)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32: // same as `rune`
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		s = strconv.FormatBool(v)
	case string:
		s = v
	case []byte:
		s = string(v)
	case time.Duration:
		s = v.String()
	case json.Number:
		s = v.String()
	default:
		s = fmt.Sprint(v)
	}
	return s
}

// IntToString int => string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString int64 => string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Uint64ToString uint64 => string
func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// Float64ToString float64 => string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float32ToString float32 => string
func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// StringToFloat64 string => float64
func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// StringToFloat32 string => float32
func StringToFloat32(s string) float32 {
	f64, _ := strconv.ParseFloat(s, 32)
	return float32(f64)
}

// StringToInt string => int
func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// StringToInt32 string => int32
func StringToInt32(s string) int32 {
	return int32(StringToInt64(s))
}

// StringToInt64 string => int64
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// StringToUint64 string => uint64
func StringToUint64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

// IntToUint int => uint
func IntToUint(i int) uint {
	return uint(i)
}

// UintToInt uint => int
func UintToInt(i uint) int {
	return int(i)
}

// JsonNumberToInt json.Number => int
func JsonNumberToInt(n json.Number) int {
	i64, _ := n.Int64()
	return int(i64)
}

// MapToJson map => json
func MapToJson(m map[string]string) (string, error) {
	b, e := json.Marshal(m)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// JsonToMap json => map
func JsonToMap(s string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Base64Encode base64 编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64Decode base64 解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// BoolToInt ...
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool ...
func IntToBool(i int) bool {
	if i > 0 {
		return true
	}
	return false
}

// ToJson ...
func ToJson(v interface{}) string {
	b, e := json.Marshal(v)
	if e != nil {
		return ""
	}
	return string(b)
}

// ToJSONBytes ...
func ToJSONBytes(v interface{}) []byte {
	b, e := json.Marshal(v)
	if e != nil {
		return nil
	}
	return b
}

// ToJSONRaw ...
func ToJSONRaw(v interface{}) json.RawMessage {
	return ToJSONBytes(v)
}

// CopyProperties 属性拷贝
func CopyProperties(source interface{}, target interface{}) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, target)
	if err != nil {
		return err
	}
	return nil
}
