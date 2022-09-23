package check

import (
	"strings"
	"unicode"
)

var (
	coefficient = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	code        = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// IsIdCard 校验一个身份证是否是合法的身份证
func IsIdCard(idCardNo string) bool {
	if len(idCardNo) != 18 {
		return false
	}

	idByte := []byte(strings.ToUpper(idCardNo))

	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(idByte[i]-byte('0')) * coefficient[i]
	}
	return code[sum%11] == idByte[17]
}

//HasHanChar 判断字符串是否包含中文字符
func HasHanChar(s string) bool {
	for _, r := range s {
		if IsHan(r) {
			return true
		}
	}
	return false
}

// IsHan http://unicode.org/Public/UNIDATA/Blocks.txt
func IsHan(r rune) bool {
	return unicode.Is(unicode.Scripts["Han"], r)
}

// IsAscii ...
func IsAscii(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}

// IsSBC 判断一个字符是否是半角字符
func IsSBC(c rune) bool {
	return IsHalfWidth(c)
}

// IsHalfWidth 判断一个字符是否是半角字符
func IsHalfWidth(c rune) bool {
	switch {
	case c <= 0x007E:
	case c == 0x00A2 || c == 0x00A3:
	case c == 0x00A5 || c == 0x00A6:
	case c == 0x00AC || c == 0x00AF:
	case c == 0x20A9:
	case c == 0x2985 || c == 0x2986:
	case c >= 0xFF61 && c <= 0xFF9F:
	case c >= 0xFFA0 && c <= 0xFFBE:
	case c >= 0xFFC2 && c <= 0xFFC7:
	case c >= 0xFFCA && c <= 0xFFCF:
	case c >= 0xFFD2 && c <= 0xFFD7:
	case c >= 0xFFDA && c <= 0xFFDC:
	case c >= 0xFFE8 && c <= 0xFFEE:
	default:
		return false
	}
	return true
}
