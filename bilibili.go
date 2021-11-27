package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const BiliHost = "https://api.live.bilibili.com"

const UserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Mobile Safari/537.36 Edg/92.0.902.67"

var reqPool = sync.Pool{New: func() interface{} {
	req, _ := http.NewRequest("GET", BiliHost, nil)
	req.Header.Set("User-Agent", UserAgent)
	return req
}}

func parseAPI(resp *http.Response) (json.RawMessage, error) {
	var biliResp BiliResp
	data, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(data, &biliResp)
	if err != nil {
		log.Println("解析API出错：", err)
		return nil, err
	}
	if biliResp.Code != 0 {
		return nil, errors.New(biliResp.Message)
		log.Panicln("API返回了错误：", biliResp.Message)
	}
	return data, nil
}

func getAPI(path, query string, v ...interface{}) json.RawMessage {
	goto start
retry:
	time.Sleep(2 * time.Second)
start:
	req := reqPool.Get().(*http.Request)
	defer reqPool.Put(req)
	req.URL.Path = path
	req.URL.RawQuery = fmt.Sprintf(query, v...)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("请求API出错：", err)
		goto retry
	}

	var biliResp BiliResp
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &biliResp)
	if err != nil {
		log.Println("解析API出错：", err)
		goto retry
	}
	if biliResp.Code != 0 {
		log.Panicln("API返回了错误：", biliResp.Message)
	}
	return biliResp.Data
}
