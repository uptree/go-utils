package stringutil

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

var (
	// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// Reverse 反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// UcFirst 首字母大写
func UcFirst(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}
	return s
}

// LcFirst 首字母小写
func LcFirst(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}
	return s
}

// CamelToSnake camel => snake 简单实现
func CamelToSnake(s string) string {
	buffer := new(bytes.Buffer)
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// SnakeToCamel snake => camel 简单实现
func SnakeToCamel(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func SnakeToSpinal(s string) string {
	return strings.Replace(s, "_", "-", -1)
}

func SpinalToSnake(s string) string {
	return strings.Replace(s, "-", "_", -1)
}

// UrlEncode 空格被编码为+，+被编码为%2B
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// UrlDecode URL解码
func UrlDecode(s string) string {
	u, _ := url.QueryUnescape(s)
	return u
}

// Substr 字符串切割
func Substr(s string, pos, length int) string {
	r := []rune(s)
	sl := len(r)
	if pos >= sl {
		return ""
	}
	idx := pos + length
	if length == 0 || idx > sl {
		idx = sl
	} else if length < 0 {
		idx = sl + length
	}

	return string(r[pos:idx])
}

// InString 判断子字符串是否存在
func InString(sub, str string) bool {
	if str != "" && strings.Contains(str, sub) {
		return true
	}
	return false
}

// RegexpReplace ...
func RegexpReplace(src, expr, repl string) (string, error) {
	reg, err := regexp.Compile(expr)
	return reg.ReplaceAllString(src, repl), err
}

// TrimSpace 去除字符串前后空格、换行等
func TrimSpace(s string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(s, trimChars)
}

// ReplaceSpace 去除字符串全部空格、换行等
func ReplaceSpace(s string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	for _, v := range trimChars {
		s = strings.ReplaceAll(s, string(v), "")
	}
	return s
}

// JoinSkipEmpty ...
func JoinSkipEmpty(sep string, ss ...string) string {
	var buf bytes.Buffer
	for _, v := range ss {
		if v == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}
