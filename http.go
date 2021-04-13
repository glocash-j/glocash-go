package glocash

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// PostForm is same as browser form submit
func PostForm(urlPath string, data map[string]string) []byte {
	var param = url.Values{}
	for k, v := range data {
		param.Set(k, v)
	}
	client := &http.Client{Timeout: 20 * time.Second}
	rsp, err := client.PostForm(urlPath, param)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	return r
}
