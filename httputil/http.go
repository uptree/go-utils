package httputil

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Deprecated: Use github.com/go-resty/resty/v2 instead.
func PostWithHeader(url string, msg []byte, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Deprecated: Use github.com/go-resty/resty/v2 instead.
func PostWithAuthorization(url, authorization string, msg []byte) (string, error) {
	headers := make(map[string]string)
	headers["Authorization"] = authorization
	headers["Content-Type"] = "application/json"
	return PostWithHeader(url, msg, headers)
}

// Deprecated: Use github.com/go-resty/resty/v2 instead.
func PostForm(url string, params url.Values) (string, error) {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// Deprecated: Use github.com/go-resty/resty/v2 instead.
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
