package httputil

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestPostWithHeader(t *testing.T) {
	apiURL := "http://localhost:8080"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body := make(map[string]string)
	body["a"] = "a"
	msg, _ := json.Marshal(body)
	_, err := PostWithHeader(apiURL, msg, headers)
	assert.NotNilf(t, err, "FAIL")
}

func TestPostWithAuthorization(t *testing.T) {
	apiURL := "http://localhost:8080"
	body := make(map[string]string)
	body["a"] = "a"
	msg, _ := json.Marshal(body)
	_, err := PostWithAuthorization(apiURL, "", msg)
	assert.NotNilf(t, err, "FAIL")
}

func TestPostForm(t *testing.T) {
	apiURL := "http://localhost:8080"
	values := url.Values{}
	values.Add("a", "a")
	_, err := PostForm(apiURL, values)
	assert.NotNilf(t, err, "FAIL")
}

func TestGet(t *testing.T) {
	apiURL := "http://localhost:8080"
	_, err := Get(apiURL)
	assert.NotNilf(t, err, "FAIL")
}
