package tools

import (
	"encoding/json"
	"fmt"
	"github.com/tss182/api-go"
	"github.com/tss182/numx"
	"time"
)

func CurlJson(api api.Api) string {
	var curlStr string
	curlStr = fmt.Sprintf("curl --request %s \\\n", api.Method)
	header := api.GetHeader()
	for i, v := range header {
		curlStr += fmt.Sprintf("-H \"%s:%s\" \\\n", i, v[0])
	}
	bodyReq, _ := json.Marshal(api.GetBody())
	if api.GetBody() != nil {
		curlStr += fmt.Sprintf("-d '%s' \\\n", bodyReq)
	}
	curlStr += fmt.Sprintf("%s", api.Url)
	return curlStr
}

func GetResp(title string, api api.Api, timeProcess time.Duration) *HistoryRequest {
	var header = map[string]interface{}{}
	for i, v := range api.GetHeader() {
		header[i] = v[0]
	}
	fmt.Println(timeProcess.Seconds())
	var respData interface{}
	_ = api.Get(&respData)
	resp := HistoryRequest{
		ID:        numx.GenerateID(),
		Title:     title,
		Url:       api.Url,
		Method:    api.Method,
		Header:    header,
		Request:   api.GetBody(),
		Response:  respData,
		Duration:  timeProcess.Seconds(),
		CreatedAt: time.Now(),
	}

	return &resp
}
