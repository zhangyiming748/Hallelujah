package Hallelujah

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/Hallelujah/log"
	"io"
	"net/http"
	"os"
)

type Img struct {
	Code   string `json:"code"`
	Acgurl string `json:"acgurl"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Size   string `json:"size"`
}

func Hallelujah() (Img, error) {
	var img Img
	b := apifox()
	err := json.Unmarshal(b, &img)
	if err != nil {
		log.Debug.Println(err)
		return Img{}, err
	} else {
		log.Info.Println("解析到结构体")
	}
	return img, nil
}
func writeAll(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Debug.Println(err)
	}
	//defer f.Close()
	n, err := f.WriteString(content)
	if err != nil {
		log.Debug.Println("写文件出错")
	} else {
		log.Info.Printf("写入%d个字节", n)
	}
}
func apifox() []byte {
	url := "https://api.jrsgslb.cn/cos/url.php?return=json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Debug.Println(err)
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		log.Debug.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Debug.Println(err)
	}
	fmt.Println(string(body))
	return body
}
