package net

import (
	"github.com/zhangyiming748/Hallelujah/log"
	"io"
	"net/http"
	"net/url"
)

func HttpGet(addHeaders map[string]string, data map[string]string, urlPath string) (body []byte, err error) {
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		log.Debug.Println(err)
	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Debug.Println(err)
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = io.ReadAll(resp.Body)
	return
}
