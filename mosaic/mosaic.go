package mosaic

import "strings"

// SafeIdCard 身份证号脱敏
func SafeIdCard(idCard string) string {
	cp := idCard
	lens := len(cp)
	return cp[0:4] + " **** **** " + cp[lens-4:]
}

// SafeMobile 手机号脱敏
func SafeMobile(mobile string) string {
	cp := mobile
	lens := len(cp)
	return cp[0:3] + "****" + cp[lens-4:]
}

// Safe 自定义字符脱敏
func Safe(cp string, front, end int) string {
	lens := len(cp)
	return cp[0:front] + strings.Repeat("*", lens-front-end) + cp[lens-end:]
}

// SafeChineseName 中文姓名脱敏
func SafeChineseName(str string) string {
	var result string
	nameRune := []rune(str)
	lens := len(nameRune)
	if lens <= 1 {
		result = "**"
	} else if lens == 2 {
		result = string(nameRune[:1]) + "*"
	} else if lens == 3 {
		result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
	} else if lens == 4 {
		result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
	} else if lens > 4 {
		result = string(nameRune[:1]) + strings.Repeat("*", lens-2) + string(nameRune[lens-1:lens])
	}
	return result
}

// SafeEmail 邮箱脱敏
func SafeEmail(str string) string {
	var result string
	res := strings.Split(str, "@")
	lens := len(res[0])
	if lens <= 1 {
		resString := "*"
		result = resString + "@" + res[1]
	} else {
		rs := []rune(str)
		res2 := string(rs[0:1])
		resString := res2 + strings.Repeat("*", lens-1)
		result = resString + "@" + res[1]
	}
	return result
}
