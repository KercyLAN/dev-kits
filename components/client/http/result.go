package http

import "net/http"

// 对HTTP客户端请求结果的结构描述
type Result struct {
	data 		[]byte			// 请求响应文包含的数据
	request 	*http.Request	// 请求体
	response 	*http.Response	// 响应体
}

// 构建一个请求结果描述实例
func newResult(data []byte, request *http.Request, response *http.Response) *Result {
	slf := &Result{
		data:data,
		request:request,
		response:response,
	}
	return slf
}