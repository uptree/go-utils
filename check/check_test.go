package check

import (
	"github.com/uptree/go-utils/assert"
	"testing"
)

func TestIsIdCard(t *testing.T) {
	assert.Equalf(t, false, IsIdCard("123456789012345678"), "FAIL")
}

func TestHasHanChar(t *testing.T) {
	assert.Equalf(t, true, HasHanChar("水蜜🍑"), "FAIL")
	assert.Equalf(t, false, HasHanChar("hello world"), "FAIL")
}

func TestIsHan(t *testing.T) {
	assert.Equalf(t, false, IsHan('🍑'), "FAIL")
	assert.Equalf(t, true, IsHan('桃'), "FAIL")
	assert.Equalf(t, false, IsHan('p'), "FAIL")
}

func TestIsAscii(t *testing.T) {
	assert.Equalf(t, false, IsAscii('🍑'), "FAIL")
	assert.Equalf(t, true, IsAscii('A'), "FAIL")
	assert.Equalf(t, true, IsAscii('0'), "FAIL")
	assert.Equalf(t, true, IsAscii('a'), "FAIL")
}
func TestIsSBC(t *testing.T) {
	assert.Equalf(t, true, IsSBC(','), "FAIL")
	assert.Equalf(t, false, IsHalfWidth('，'), "FAIL")
}

func TestIsHalfWidth(t *testing.T) {
	assert.Equalf(t, true, IsHalfWidth(','), "FAIL")
	assert.Equalf(t, false, IsHalfWidth('，'), "FAIL")
	assert.Equalf(t, true, IsHalfWidth('<'), "FAIL")
	assert.Equalf(t, false, IsHalfWidth('《'), "FAIL")
	assert.Equalf(t, true, IsHalfWidth('$'), "FAIL")
	// 半角: '¥' U+00A5 全角: '￥' U+FFE5 建议使用半角
	assert.Equalf(t, true, IsHalfWidth('¥'), "FAIL")
	assert.Equalf(t, false, IsHalfWidth('你'), "FAIL")
	assert.Equalf(t, true, IsHalfWidth('n'), "FAIL")
}
