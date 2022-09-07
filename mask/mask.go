package mask

import "strings"

// IdCard 身份证号脱敏
func IdCard(s string) string {
	return s[0:4] + " **** **** " + s[len(s)-4:]
}

// IdCardStrict 严格身份证号脱敏
func IdCardStrict(s string) string {
	return s[0:1] + "*** **** **** ***" + s[len(s)-1:]
}

// Mobile 手机号脱敏
func Mobile(s string) string {
	return s[0:3] + "****" + s[len(s)-4:]
}

// LastFour 后四位
func LastFour(s string) string {
	l := len(s)
	if l <= 4 {
		return s
	}
	return s[l-4:]
}

// Mask 自定义字符脱敏
func Mask(cp string, front, end int) string {
	l := len(cp)
	return cp[0:front] + strings.Repeat("*", l-front-end) + cp[l-end:]
}

// ChineseName 中文姓名脱敏
func ChineseName(s string) string {
	r := []rune(s)
	l := len(r)
	if l == 2 {
		return "*" + string(r[1:])
	} else if l == 3 {
		return "*" + string(r[l-2:])
	} else if l == 4 {
		return "**" + string(r[l-2:])
	} else if l > 4 {
		return string(r[:1]) + strings.Repeat("*", l-3) + string(r[l-2:])
	}
	return "**"
}

// Email 邮箱脱敏
func Email(s string) string {
	ss := strings.Split(s, "@")
	l := len(ss[0])
	if l <= 1 {
		return "*@" + ss[1]
	}
	r := []rune(s)
	return string(r[0:1]) + strings.Repeat("*", l-1) + "@" + ss[1]
}
