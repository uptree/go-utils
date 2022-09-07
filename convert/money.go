package convert

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/uptree/go-utils/stringutil"
)

var (
	// ChineseSliceUnit ...
	ChineseSliceUnit = []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "圆", "角", "分"}
	// ChineseDigitUpper ...
	ChineseDigitUpper = map[string]string{
		"0": "零",
		"1": "壹",
		"2": "贰",
		"3": "叁",
		"4": "肆",
		"5": "伍",
		"6": "陆",
		"7": "柒",
		"8": "捌",
		"9": "玖",
	}
	// ChineseUpperDigit ...
	ChineseUpperDigit = map[int32]int8{
		'零': 0,
		'壹': 1,
		'贰': 2,
		'两': 2,
		'叁': 3,
		'肆': 4,
		'伍': 5,
		'陆': 6,
		'柒': 7,
		'捌': 8,
		'玖': 9,
	}
	// chineseRegexp ...
	chineseRegexp = [][]string{
		{`零角零分$`, "整"},
		{`零角`, "零"},
		{`零分$`, "整"},
		{`零[仟佰拾]`, "零"},
		{`零{2,}`, "零"},
		{`零亿`, "亿"},
		{`零万`, "万"},
		{`亿万`, "亿"},
		{`零*圆`, "圆"},
		{`亿零{0, 3}万`, "^圆"},
		{`零圆`, "零"},
	}
)

// FenToYuan 货币分转元
func FenToYuan(fen int) float64 {
	return float64(fen) / 100
}

//YuanToFen 货币元转分
func YuanToFen(yuan float64) int {
	return int(math.Round(yuan * 100))
}

// FenToStringYuan 分转字符串元
func FenToStringYuan(fen int) string {
	s := IntToString(fen)
	if fen < 10 {
		return "0.0" + s
	}
	if fen < 100 {
		return "0." + s
	}
	return s[0:len(s)-2] + "." + s[len(s)-2:]
}

// StringYuanToFen 字符串元转分 保留小数2位 超过两位会被舍弃
func StringYuanToFen(yuan string) int {
	arr := strings.Split(yuan, ".")
	fen := StringToInt(arr[0]) * 100
	if len(arr) > 1 {
		if len(arr[1]) >= 2 {
			fen += StringToInt(arr[1][0:2])
		} else {
			fen += StringToInt(arr[1][0:1]) * 10
		}
	}
	return fen
}

// FenToChineseYuan 货币数字分转中文元
func FenToChineseYuan(fen int) (string, error) {
	sFen := strconv.Itoa(fen)
	if fen <= 0 || len(sFen) > len(ChineseSliceUnit) {
		return "", errors.New("out of range")
	}
	var (
		cYuan string
		err   error
	)
	s := ChineseSliceUnit[len(ChineseSliceUnit)-len(sFen):]
	for k, v := range sFen[:] {
		cYuan = cYuan + ChineseDigitUpper[string(v)] + s[k]
	}
	for _, v := range chineseRegexp {
		cYuan, err = stringutil.RegexpReplace(cYuan, v[0], v[1])
	}
	if err != nil {
		return "", err
	}
	return cYuan, err
}

// ChineseYunToFen 大写元转数字分
func ChineseYunToFen(yuan string) int {
	interValue := make([]byte, 16)
	floatValue := make([]byte, 2)
	s := stringutil.Reverse(yuan)
	var interPlace = 0
	var bigPlace = 0
	var floatPlace = 0
	for _, row := range s {
		if row == 0 || row == '零' {
			continue
		}
		//确定位
		switch row {
		case '圆':
			interPlace = 0
			bigPlace = 0
			//标识不在是小数部份
			floatPlace = -1
		case '拾':
			interPlace = 1
		case '佰':
			interPlace = 2
		case '仟':
			interPlace = 3
		case '万':
			interPlace = 0
			bigPlace = 4
		case '亿':
			interPlace = 0
			bigPlace = 8
		case '角':
			floatPlace = 0
		case '分':
			floatPlace = 1
		default:
			if floatPlace == -1 {
				//整数部份
				interValue[interPlace+bigPlace] = byte(ChineseUpperDigit[row])
			} else {
				//小数部份
				floatValue[floatPlace] = byte(ChineseUpperDigit[row])
			}
		}
	}

	//从写入的位置整理成最终的数字
	var num float64 = 0
	for i := 0; i < len(interValue); i++ {
		num += float64(interValue[i]) * math.Pow(10, float64(i))
	}

	for i := 0; i < len(floatValue); i++ {
		num += float64(floatValue[i]) * math.Pow(1.0/10.0, float64(i+1))
	}
	return YuanToFen(num)
}
