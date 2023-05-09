package utils

import (
	"io"
	"net/http"
)

// HttpGet 发起get请求， 并返回string类型的body
func HttpGet(url string) (body string, err error) {
	// http.get发起请求， 并返回string类型的body
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取body
	body, err = ReadBody(res.Body)
	if err != nil {
		return "", err
	}
	return body, nil
}

// ReadBody 读取body
func ReadBody(body io.ReadCloser) (string, error) {
	// 读取body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}
	// 转换为string类型
	bodyString := string(bodyBytes)
	return bodyString, nil
}
