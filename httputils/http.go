package httputils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostWithAuthorization(url, authorization string, msg []byte) (string, error) {
	headers := make(map[string]string)
	headers["Authorization"] = authorization
	headers["Content-Type"] = "application/json"
	return PostWithHeader(url, msg, headers)
}
