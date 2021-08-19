package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	//resp, _ := http.Get("http://127.0.0.1:8000/go?page=1&pageSize=10")
	//// 200 OK
	//fmt.Println(resp.Status)
	//fmt.Println(resp.Header)
	urlObj, _ := url.Parse("http://127.0.0.1:8000/go/")
	data := url.Values{} // url encode
	data.Set("page", "1")
	data.Set("pageSize", "10")
	urlObj.RawQuery = data.Encode() // URL Encode之后的URL
	fmt.Println(urlObj.String(), urlObj.RawQuery)
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	//resp, _ := http.DefaultClient.Do(req)
	// 禁用KeepAlive的client 拉取不是特别频繁，则禁用长连接
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp,_ := client.Do(req)
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
