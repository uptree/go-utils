package mask

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestMask(t *testing.T) {
	cp := "13844239123"
	n := len(cp)
	s := Mask(cp, 1, 2)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Equalf(t, "1********23", s, "FAIL")
}

func TestBefore(t *testing.T) {
	cp := "13844239123"
	n := len(cp)
	s := Before(cp, 1)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Equalf(t, "1**********", s, "FAIL")
}

func TestAfter(t *testing.T) {
	cp := "13844239123"
	n := len(cp)
	s := After(cp, 1)
	assert.Lenf(t, s, n, "长度:%d", n)
	assert.Equalf(t, "**********3", s, "FAIL")
}

func TestFirst(t *testing.T) {
	assert.Equalf(t, "123", First("123", 4), "FAIL")
	assert.Equalf(t, "1380", First("13800001234", 4), "FAIL")
	assert.Equalf(t, "", First("460220202012210027", 0), "FAIL")
}

func TestLast(t *testing.T) {
	assert.Equalf(t, "123", Last("123", 4), "FAIL")
	assert.Equalf(t, "1234", Last("13800001234", 4), "FAIL")
	assert.Equalf(t, "210027", Last("460220202012210027", 6), "FAIL")
}

func TestLastFour(t *testing.T) {
	assert.Equalf(t, "123", LastFour("123"), "FAIL")
	assert.Equalf(t, "1234", LastFour("13800001234"), "FAIL")
	assert.Equalf(t, "0027", LastFour("460220202012210027"), "FAIL")
}

func TestIdCard(t *testing.T) {
	assert.Equalf(t, "4602 **** **** 0027", IdCard("460220202012210027"), "FAIL")
}

func TestIdCardStrict(t *testing.T) {
	assert.Equalf(t, "4*** **** **** ***7", IdCardStrict("460220202012210027"), "FAIL")
}

func TestMobile(t *testing.T) {
	assert.Equalf(t, "138****1234", Mobile("13800001234"), "FAIL")
}

func TestChineseName(t *testing.T) {
	assert.Equalf(t, "**", ChineseName("郑"), "FAIL")
	assert.Equalf(t, "*爽", ChineseName("郑爽"), "FAIL")
	assert.Equalf(t, "*杰伦", ChineseName("周杰伦"), "FAIL")
	assert.Equalf(t, "**娜娜", ChineseName("欧阳娜娜"), "FAIL")
	assert.Equalf(t, "斯**日乐", ChineseName("斯琴格日乐"), "FAIL")
	assert.Equalf(t, "阿********西提", ChineseName("阿不来提 阿卜杜热西提"), "FAIL")
	assert.Equalf(t, "C***尔多", ChineseName("C 罗纳尔多"), "FAIL")

}

func TestEmail(t *testing.T) {
	assert.Equalf(t, "*@g.com", Email("1@g.com"), "FAIL")
	assert.Equalf(t, "1*@tsinghua.org.cn", Email("10@tsinghua.org.cn"), "FAIL")
	assert.Equalf(t, "1**@163.com", Email("100@163.com"), "FAIL")
	assert.Equalf(t, "1***@foxmail.com", Email("1000@foxmail.com"), "FAIL")
	assert.Equalf(t, "1****@sina.com.cn", Email("10000@sina.com.cn"), "FAIL")
	assert.Equalf(t, "1**********@gmail.com", Email("13312345678@gmail.com"), "FAIL")
}
