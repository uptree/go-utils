package mosaic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafe(t *testing.T) {
	cp := "13844239123"
	n := len(cp)
	s := Safe(cp, 1, 2)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Equalf(t, "1********23", s, "failed")
}

func TestSafeIdCard(t *testing.T) {
	assert.Equalf(t, "4602 **** **** 0027", SafeIdCard("460220202012210027"), "failed")
}

func TestSafeMobile(t *testing.T) {
	assert.Equalf(t, "138****1234", SafeMobile("13800001234"), "failed")
}

func TestSafeChineseName(t *testing.T) {
	assert.Equalf(t, "**", SafeChineseName("郑"), "failed")
	assert.Equalf(t, "郑*", SafeChineseName("郑爽"), "failed")
	assert.Equalf(t, "周*伦", SafeChineseName("周杰伦"), "failed")
	assert.Equalf(t, "欧**娜", SafeChineseName("欧阳娜娜"), "failed")
	assert.Equalf(t, "斯***乐", SafeChineseName("斯琴格日乐"), "failed")
	assert.Equalf(t, "阿*********提", SafeChineseName("阿不来提 阿卜杜热西提"), "failed")
	assert.Equalf(t, "C****多", SafeChineseName("C 罗纳尔多"), "failed")

}

func TestSafeEmail(t *testing.T) {
	assert.Equalf(t, "*@g.com", SafeEmail("1@g.com"), "failed")
	assert.Equalf(t, "1*@tsinghua.org.cn", SafeEmail("10@tsinghua.org.cn"), "failed")
	assert.Equalf(t, "1**@163.com", SafeEmail("100@163.com"), "failed")
	assert.Equalf(t, "1***@foxmail.com", SafeEmail("1000@foxmail.com"), "failed")
	assert.Equalf(t, "1****@sina.com.cn", SafeEmail("10000@sina.com.cn"), "failed")
	assert.Equalf(t, "1**********@gmail.com", SafeEmail("13312345678@gmail.com"), "failed")
}
