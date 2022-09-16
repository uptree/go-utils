package convert

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestFenToYuan(t *testing.T) {
	assert.Equalf(t, 19.90, FenToYuan(1990), "FAIL")
	t.Logf("FenToYuan: %f", FenToYuan(10000000000001)) // 精度异常
	t.Logf("FenToYuan: %.2f", FenToYuan(10000000000001))
}

func TestYuanToFen(t *testing.T) {
	assert.Equalf(t, 1990, YuanToFen(19.90), "FAIL")
	assert.Equalf(t, 104911, YuanToFen(1049.11), "FAIL")
}

func TestFenToStringYuan(t *testing.T) {
	assert.Equalf(t, "0.01", FenToStringYuan(1), "FAIL")
	assert.Equalf(t, "0.10", FenToStringYuan(10), "FAIL")
	assert.Equalf(t, "1.01", FenToStringYuan(101), "FAIL")
	assert.Equalf(t, "10.01", FenToStringYuan(1001), "FAIL")
	assert.Equalf(t, "10.10", FenToStringYuan(1010), "FAIL")
	assert.Equalf(t, "10.00", FenToStringYuan(1000), "FAIL")
	assert.Equalf(t, "1000000000000000.00", FenToStringYuan(100000000000000000), "FAIL")
}

func TestStringYuanToFen(t *testing.T) {
	assert.Equalf(t, 1990, StringYuanToFen("19.90"), "FAIL")
	assert.Equalf(t, 1990, StringYuanToFen("19.9091"), "FAIL")
	assert.Equalf(t, 1990, StringYuanToFen("19.9"), "FAIL")
	assert.Equalf(t, 1900, StringYuanToFen("19"), "FAIL")
	assert.Equalf(t, -100, StringYuanToFen("1000000000000000000000000"), "FAIL")
}

func TestFenToChineseYuan(t *testing.T) {

	s, e := FenToChineseYuan(100000000000000) //out of range
	assert.Equalf(t, "out of range", e.Error(), "FAIL")
	s, e = FenToChineseYuan(10101010101101)
	assert.Equalf(t, "壹仟零壹拾亿壹仟零壹拾万壹仟零壹拾壹圆零壹分", s, "FAIL")
	s, e = FenToChineseYuan(123000000111)
	assert.Equalf(t, "壹拾贰亿叁仟万零壹圆壹角壹分", s, "FAIL")
	s, e = FenToChineseYuan(2200)
	assert.Equalf(t, "贰拾贰圆整", s, "FAIL")
	s, e = FenToChineseYuan(5000000)
	assert.Equalf(t, "伍万圆整", s, "FAIL")
	s, e = FenToChineseYuan(100)
	assert.Equalf(t, "壹圆整", s, "FAIL")
	s, e = FenToChineseYuan(10101010101101)
	assert.Equalf(t, "壹仟零壹拾亿壹仟零壹拾万壹仟零壹拾壹圆零壹分", s, "FAIL")
	s, e = FenToChineseYuan(10000000000001)
	assert.Equalf(t, "壹仟亿圆零壹分", s, "FAIL")
	s, e = FenToChineseYuan(1990)
	assert.Equalf(t, "壹拾玖圆玖角整", s, "FAIL")
}

func TestChineseYunToFen(t *testing.T) {
	assert.Equalf(t, 10000000000001, ChineseYunToFen("壹仟亿圆零壹分"), "FAIL")
	assert.Equalf(t, 1990, ChineseYunToFen("壹拾玖圆玖角整"), "FAIL")
	assert.Equalf(t, 200000, ChineseYunToFen("两仟圆整"), "FAIL")
}
