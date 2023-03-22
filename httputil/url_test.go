package httputil

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestMapToRawQuery(t *testing.T) {
	params := make(map[string]string)
	params["Appid"] = "10001"
	params["Version"] = "20220914"
	params["Timestamp"] = "1663144810"
	params["Nonce"] = "873926"
	query := "Appid=10001&Nonce=873926&Timestamp=1663144810&Version=20220914"
	assert.Equalf(t, query, MapToRawQuery(params), "FAIL")
}

func TestRawQueryToMap(t *testing.T) {
	params := make(map[string]string)
	params["Appid"] = "10001"
	params["Version"] = "20220914"
	params["Timestamp"] = "1663144810"
	params["Nonce"] = "873926"
	query := "Appid=10001&Nonce=873926&Timestamp=1663144810&Version=20220914"
	assert.Equalf(t, params, RawQueryToMap(query), "FAIL")
}

func TestRawURLGetScheme(t *testing.T) {
	assert.Equalf(t, "https", RawURLGetScheme("https://localhost:8080/index"), "FAIL")
	assert.Equalf(t, "http", RawURLGetScheme("Http://localhost:8080/index"), "FAIL")
	assert.Equalf(t, "scheme", RawURLGetScheme("Scheme://localhost:8080/index"), "FAIL")
	assert.Emptyf(t, RawURLGetScheme("//localhost:8080/index"), "FAIL")
	assert.Emptyf(t, RawURLGetScheme("://localhost:8080/index"), "FAIL")
}

func TestRawURLGetHost(t *testing.T) {
	assert.Equalf(t,
		"localhost:8080", RawURLGetHost("https://localhost:8080/index?b=1&a=2"), "FAIL")
	assert.Equalf(t,
		"localhost:8080", RawURLGetHost("xx://localhost:8080/index?b=1&a=2"), "FAIL")
}

func TestRawURLGetParams(t *testing.T) {
	rawURL := "https://localhost/index?b=1&a=2&A=1"
	m := make(map[string][]string)
	m["a"] = []string{"2"}
	m["A"] = []string{"1"}
	m["b"] = []string{"1"}
	params, _ := RawURLGetParams(rawURL)
	assert.Equalf(t, m, params, "FAIL")
}

func TestRawURLAddParams(t *testing.T) {
	rawURL := "https://localhost/index?b=1&a=2"
	params := make(map[string]string)
	params["a"] = "3"
	// 新增，参数会被重新排序
	assert.Equalf(t,
		"https://localhost/index?a=2&a=3&b=1", RawURLAddParams(rawURL, params), "FAIL")
}

func TestRawURLSetParams(t *testing.T) {
	rawURL := "https://localhost/index?b=1&a=2"
	params := make(map[string]string)
	params["a"] = "3"
	// 覆盖，参数会被重新排序
	assert.Equalf(t,
		"https://localhost/index?a=3&b=1", RawURLSetParams(rawURL, params), "FAIL")
}

func TestRawURLDelParams(t *testing.T) {
	rawURL := "https://localhost/index?b=1&a=2"
	assert.Equalf(t,
		"https://localhost/index", RawURLDelParams(rawURL, []string{"a", "b"}), "FAIL")
}

func TestRawURLToHttps(t *testing.T) {
	rawURL := "http://localhost//index?b=1&a=2"
	newURL := "https://localhost//index?b=1&a=2"
	assert.Equalf(t, newURL, RawURLToHttps(rawURL), "FAIL")
}

func TestTrimScheme(t *testing.T) {
	assert.Equalf(t, "localhost//index?b=1&a=2",
		TrimScheme("http://localhost//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, "localhost:8001//index?b=1&a=2",
		TrimScheme("cc://localhost:8001//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, "localhost//index?b=1&a=2",
		TrimScheme("ftp://localhost//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, "localhost//index?b=1&a=2",
		TrimScheme("https://localhost//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, "localhost//index?b=1&a=2",
		TrimScheme("//localhost//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, "/localhost//index?b=1&a=2",
		TrimScheme("/localhost//index?b=1&a=2"), "FAIL")
	assert.Equalf(t, ":/localhost//index?b=1&a=2",
		TrimScheme(":/localhost//index?b=1&a=2"), "FAIL")
}
