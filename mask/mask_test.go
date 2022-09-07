package mask

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMask(t *testing.T) {
	cp := "13844239123"
	n := len(cp)
	s := Mask(cp, 1, 2)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Equalf(t, "1********23", s, "failed")
}

func TestIdCard(t *testing.T) {
	assert.Equalf(t, "4602 **** **** 0027", IdCard("460220202012210027"), "failed")
}

func TestIdCardStrict(t *testing.T) {
	assert.Equalf(t, "4*** **** **** ***7", IdCardStrict("460220202012210027"), "failed")
}

func TestMobile(t *testing.T) {
	assert.Equalf(t, "138****1234", Mobile("13800001234"), "failed")
}

func TestLastFour(t *testing.T) {
	assert.Equalf(t, "123", LastFour("123"), "failed")
	assert.Equalf(t, "1234", LastFour("13800001234"), "failed")
	assert.Equalf(t, "0027", LastFour("460220202012210027"), "failed")
}

func TestChineseName(t *testing.T) {
	assert.Equalf(t, "**", ChineseName("郑"), "failed")
	assert.Equalf(t, "*爽", ChineseName("郑爽"), "failed")
	assert.Equalf(t, "*杰伦", ChineseName("周杰伦"), "failed")
	assert.Equalf(t, "**娜娜", ChineseName("欧阳娜娜"), "failed")
	assert.Equalf(t, "斯**日乐", ChineseName("斯琴格日乐"), "failed")
	assert.Equalf(t, "阿********西提", ChineseName("阿不来提 阿卜杜热西提"), "failed")
	assert.Equalf(t, "C***尔多", ChineseName("C 罗纳尔多"), "failed")

}

func TestEmail(t *testing.T) {
	assert.Equalf(t, "*@g.com", Email("1@g.com"), "failed")
	assert.Equalf(t, "1*@tsinghua.org.cn", Email("10@tsinghua.org.cn"), "failed")
	assert.Equalf(t, "1**@163.com", Email("100@163.com"), "failed")
	assert.Equalf(t, "1***@foxmail.com", Email("1000@foxmail.com"), "failed")
	assert.Equalf(t, "1****@sina.com.cn", Email("10000@sina.com.cn"), "failed")
	assert.Equalf(t, "1**********@gmail.com", Email("13312345678@gmail.com"), "failed")
}
