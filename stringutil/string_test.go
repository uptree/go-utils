package stringutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equalf(t, "olleh", Reverse("hello"), "failed")
}
func TestUcFirst(t *testing.T) {
	assert.Equalf(t, "Hello", UcFirst("hello"), "failed")
	assert.Equalf(t, "Hello", UcFirst("Hello"), "failed")
	assert.Equalf(t, "Hello_world", UcFirst("hello_world"), "failed")
	assert.Equalf(t, "HelloWorld", UcFirst("helloWorld"), "failed")
	// 对比
	assert.Equalf(t, "Hello world", UcFirst("hello world"), "failed")
	assert.Equalf(t, "Hello World", strings.Title("hello world"), "failed")
}

func TestLcFirst(t *testing.T) {
	assert.Equalf(t, "hello", LcFirst("Hello"), "failed")
	assert.Equalf(t, "hELLO", LcFirst("HELLO"), "failed")
	assert.Equalf(t, "hELLO WORLD", LcFirst("HELLO WORLD"), "failed")
}

func TestCamelToSnake(t *testing.T) {
	assert.Equalf(t, "test_camel_to_snake", CamelToSnake("TestCamelToSnake"), "failed")
	assert.Equalf(t, "test_camel_to_snake", CamelToSnake("test_camel_to_snake"), "failed")
	// 异常  需要预处理
	assert.Equalf(t, "t_e_s_t__c_a_m_e_l__t_o__s_n_a_k_e",
		CamelToSnake("TEST_CAMEL_TO_SNAKE"), "failed")
}

func TestSnakeToCamel(t *testing.T) {
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("TestSnakeToCamel"), "failed")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("Test_SnakeToCamel"), "failed")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("testSnakeToCamel"), "failed")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("test_snake_to_camel"), "failed")
	// 异常  需要预处理
	assert.Equalf(t, "TESTSNAKETOCAMEL", SnakeToCamel("TEST_SNAKE_TO_CAMEL"), "failed")
	assert.Equalf(t, "TESTSnakeTOCAMEL", SnakeToCamel("TEST_snake_TO_CAMEL"), "failed")
}

func TestSnakeToSpinal(t *testing.T) {
	assert.Equalf(t, "test-snake-to-spinal", SnakeToSpinal("test_snake_to_spinal"), "failed")
}

func TestSpinalToSnake(t *testing.T) {
	assert.Equalf(t, "test_spinal_to_snake", SpinalToSnake("test-spinal-to-snake"), "failed")
}

func TestUrlEncode(t *testing.T) {
	assert.Equalf(t, "+%E5%A4%AA%E9%98%B3%E5%BD%93%E7%A9%BA%E7%85%A7%2C%E8%8A%B1%E5%84%BF%E5%AF%B9%E6%88%91%E7%AC%91%E3%80%82",
		UrlEncode(" 太阳当空照,花儿对我笑。"), "failed")
	assert.Equalf(t, "1+%2B+1+%3D+2", UrlEncode("1 + 1 = 2"), "failed")
}

func TestUrlDecode(t *testing.T) {
	assert.Equalf(t, " 太阳当空照,花儿对我笑。 ", UrlDecode(" 太阳当空照,花儿对我笑。+"), "failed")
	assert.Equalf(t, " 太阳当空照,花儿对我笑。",
		UrlDecode("+%E5%A4%AA%E9%98%B3%E5%BD%93%E7%A9%BA%E7%85%A7%2C%E8%8A%B1%E5%84%BF%E5%AF%B9%E6%88%91%E7%AC%91%E3%80%82"), "failed")
	assert.Equalf(t, "1   1 = 2", UrlDecode("1 + 1 = 2"), "failed")
	assert.Equalf(t, "1 + 1 = 2", UrlDecode("1+%2B+1+%3D+2"), "failed")
}

func TestSubstr(t *testing.T) {
	cnStr := " 太阳当空照,花儿对我笑。"
	assert.Equalf(t, "太阳当空照,花儿对我笑", Substr(cnStr, 1, -1), "failed")
	str := "abcdefg"
	assert.Equalf(t, "abcd", Substr(str, 0, 4), "failed")
	assert.Equalf(t, "abc", Substr(str, 0, -4), "failed")
	assert.Equalf(t, "c", Substr(str, 2, -4), "failed")
}

func TestInString(t *testing.T) {
	assert.Equalf(t, true, InString("花儿", " 太阳当空照,花儿对我笑。"), "failed")
	assert.Equalf(t, true, InString("DE", " ABCDEDFH"), "failed")
	assert.Equalf(t, false, InString("DEDE", " ABCDEDFH"), "failed")
}
