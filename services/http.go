package services

import (
	// "bytes"
    "encoding/json"
    // "io/ioutil"
    // "net/http"
	"time"
	"fmt"

	"github.com/valyala/fasthttp"
)

func Post(url string, data interface{}, contentType string) (res string, erro error){
    req := &fasthttp.Request{}
    req.SetRequestURI(url)
    jsonStr, _ := json.Marshal(data)
    // requestBody := []byte(`{"request":"test"}`)
    req.SetBody([]byte(jsonStr))

    // 默认是application/x-www-form-urlencoded
    req.Header.SetContentType(contentType)
    req.Header.SetMethod("POST")

    resp := &fasthttp.Response{}

    client := &fasthttp.Client{}
    if err := client.Do(req, resp);err != nil {		
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("请求失败:", err.Error())
		erro = err
        return
    }


	// result, _ := ioutil.ReadAll(resp.Body())
	res = string(resp.Body())
    return 
}