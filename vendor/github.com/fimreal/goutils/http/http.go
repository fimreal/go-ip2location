package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	return string(res), err
}

// 似乎只支持 https，没找到问题原因，慎用
func HttpBufGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return result, err
		}
		// 拼接每次buf中读到的数据，到result中，返回
		result += string(buf[:n])
	}
	return
}

/* SimplePost 简单 post 调用，内部做了 header 处理和返回错误判断
用法举例：
url := "http://example.com"
headers := make(map[string]string)
headers["Content-Type"] = "application/json;charset=utf-8"

resp, err := http.SimplePost(url, []byte("这里是 post 的数据"), headers)
*/
func SimplePost(url string, data []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// 省去判断 err ？
	return body, err
}

func HttpBasicAuth(url string, username string, password string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
