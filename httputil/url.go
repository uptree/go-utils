package httputil

import (
	"net/url"
	"strings"

	"github.com/uptree/go-utils/convert"
)

// Scheme 传输协议
type Scheme string

const (
	// HTTP 表示HTTP明文传输协议
	HTTP Scheme = "http"
	// HTTPS 表示HTTPS加密传输协议
	HTTPS Scheme = "https"
)

// MapToRawQuery Map => Raw Query  字典序
func MapToRawQuery(params map[string]interface{}) string {
	queries := url.Values{}
	for k, v := range params {
		// 过滤空值
		if v != "" {
			queries.Add(k, convert.AnyToString(v))
		}
	}
	return queries.Encode()
}

// RawQueryToMap Raw Query => Map
func RawQueryToMap(rawQuery string) map[string]interface{} {
	queries, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil
	}
	params := make(map[string]interface{})
	for k, v := range queries {
		if len(v) > 0 {
			params[k] = v[0] // 同名只取其一
		}
	}
	return params
}

// RawURLGetHost 获取URL的域名+端口
func RawURLGetHost(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Host
}

// RawURLGetScheme 获取URL的协议
func RawURLGetScheme(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Scheme
}

// RawURLGetParams 获取URL参数Map
func RawURLGetParams(rawURL string) (map[string][]string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return u.Query(), nil
}

// RawURLAddParams 增加参数，不会覆盖，会重新排序
func RawURLAddParams(rawURL string, params map[string]string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	m := u.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	u.RawQuery = m.Encode()
	return u.String()
}

// RawURLSetParams 增加或修改参数，会覆盖，会重新排序
func RawURLSetParams(rawURL string, params map[string]string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	m := u.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	u.RawQuery = m.Encode()
	return u.String()
}

// RawURLDelParams 删除参数
func RawURLDelParams(rawURL string, keys []string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	m := u.Query()
	for _, v := range keys {
		m.Del(v)
	}
	u.RawQuery = m.Encode()
	return u.String()
}

// RawURLToHttps URL强制转换为https
func RawURLToHttps(rawURL string) string {
	if strings.Index(rawURL, "http:") == 0 {
		return strings.Replace(rawURL, "http:", "https:", 1)
	}
	if strings.Index(rawURL, "//") == 0 {
		return "https:" + rawURL
	}
	return rawURL
}

// TrimScheme 过滤URL协议
func TrimScheme(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	return u.Host + u.RequestURI()
}
