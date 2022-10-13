package stringutil

import (
	"strings"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestReverse(t *testing.T) {
	assert.Equalf(t, "olleh", Reverse("hello"), "FAIL")
}
func TestUcFirst(t *testing.T) {
	assert.Equalf(t, "Hello", UcFirst("hello"), "FAIL")
	assert.Equalf(t, "Hello", UcFirst("Hello"), "FAIL")
	assert.Equalf(t, "Hello_world", UcFirst("hello_world"), "FAIL")
	assert.Equalf(t, "HelloWorld", UcFirst("helloWorld"), "FAIL")
	// 对比
	assert.Equalf(t, "Hello world", UcFirst("hello world"), "FAIL")
	assert.Equalf(t, "Hello World", strings.Title("hello world"), "FAIL")
}

func TestLcFirst(t *testing.T) {
	assert.Equalf(t, "hello", LcFirst("Hello"), "FAIL")
	assert.Equalf(t, "hELLO", LcFirst("HELLO"), "FAIL")
	assert.Equalf(t, "hELLO WORLD", LcFirst("HELLO WORLD"), "FAIL")
}

func TestCamelToSnake(t *testing.T) {
	assert.Equalf(t, "test_camel_to_snake", CamelToSnake("TestCamelToSnake"), "FAIL")
	assert.Equalf(t, "test_camel_to_snake", CamelToSnake("test_camel_to_snake"), "FAIL")
	// 异常  需要预处理
	assert.Equalf(t, "t_e_s_t__c_a_m_e_l__t_o__s_n_a_k_e",
		CamelToSnake("TEST_CAMEL_TO_SNAKE"), "FAIL")
}

func TestSnakeToCamel(t *testing.T) {
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("TestSnakeToCamel"), "FAIL")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("Test_SnakeToCamel"), "FAIL")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("testSnakeToCamel"), "FAIL")
	assert.Equalf(t, "TestSnakeToCamel", SnakeToCamel("test_snake_to_camel"), "FAIL")
	// 异常  需要预处理
	assert.Equalf(t, "TESTSNAKETOCAMEL", SnakeToCamel("TEST_SNAKE_TO_CAMEL"), "FAIL")
	assert.Equalf(t, "TESTSnakeTOCAMEL", SnakeToCamel("TEST_snake_TO_CAMEL"), "FAIL")
}

func TestSnakeToSpinal(t *testing.T) {
	assert.Equalf(t, "test-snake-to-spinal", SnakeToSpinal("test_snake_to_spinal"), "FAIL")
}

func TestSpinalToSnake(t *testing.T) {
	assert.Equalf(t, "test_spinal_to_snake", SpinalToSnake("test-spinal-to-snake"), "FAIL")
}

func TestUrlEncode(t *testing.T) {
	assert.Equalf(t, "+%E5%A4%AA%E9%98%B3%E5%BD%93%E7%A9%BA%E7%85%A7%2C%E8%8A%B1%E5%84%BF%E5%AF%B9%E6%88%91%E7%AC%91%E3%80%82",
		UrlEncode(" 太阳当空照,花儿对我笑。"), "FAIL")
	assert.Equalf(t, "1+%2B+1+%3D+2", UrlEncode("1 + 1 = 2"), "FAIL")
}

func TestUrlDecode(t *testing.T) {
	assert.Equalf(t, " 太阳当空照,花儿对我笑。 ", UrlDecode(" 太阳当空照,花儿对我笑。+"), "FAIL")
	assert.Equalf(t, " 太阳当空照,花儿对我笑。",
		UrlDecode("+%E5%A4%AA%E9%98%B3%E5%BD%93%E7%A9%BA%E7%85%A7%2C%E8%8A%B1%E5%84%BF%E5%AF%B9%E6%88%91%E7%AC%91%E3%80%82"), "FAIL")
	assert.Equalf(t, "1   1 = 2", UrlDecode("1 + 1 = 2"), "FAIL")
	assert.Equalf(t, "1 + 1 = 2", UrlDecode("1+%2B+1+%3D+2"), "FAIL")
}

func TestSubstr(t *testing.T) {
	cnStr := " 太阳当空照,花儿对我笑。"
	assert.Equalf(t, "太阳当空照,花儿对我笑", Substr(cnStr, 1, -1), "FAIL")
	str := "abcdefg"
	assert.Equalf(t, "abcd", Substr(str, 0, 4), "FAIL")
	assert.Equalf(t, "abc", Substr(str, 0, -4), "FAIL")
	assert.Equalf(t, "c", Substr(str, 2, -4), "FAIL")
}

func TestInString(t *testing.T) {
	assert.Equalf(t, true, InString("花儿", " 太阳当空照,花儿对我笑。"), "FAIL")
	assert.Equalf(t, true, InString("DE", " ABCDEDFH"), "FAIL")
	assert.Equalf(t, false, InString("DEDE", " ABCDEDFH"), "FAIL")
}

func TestRegexpReplace(t *testing.T) {
	s, _ := RegexpReplace("1000001", `0{2,}`, "0")
	assert.Equalf(t, "101", s, "FAIL")
}

func TestTrimSpace(t *testing.T) {
	assert.Equalf(t, "hello world", TrimSpace(" hello world \n"), "FAIL")
}

func TestReplaceSpace(t *testing.T) {
	assert.Equalf(t, "helloworld", ReplaceSpace("he llo wo\r\t r ld \n"), "FAIL")
}

func TestJoinSkipEmpty(t *testing.T) {
	assert.Equalf(t, "a|b|c|d", JoinSkipEmpty("|", []string{"a", "b", "c", "", "d"}...), "FAIL")
}
