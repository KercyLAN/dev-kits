package http

import "net/http"

// 网络请求客户端结构描述
type client struct {
	*http.Client							// 原生HTTP客户端
	todo 			chan *http.Request		// 待办请求
}

// 执行请求
func (slf *client) exec(request *http.Request) {

}

// 添加待办请求
func (slf *client) add(request *http.Request) {
	slf.todo <- request
}


// 运行客户端
func (slf *client) run(idle chan <- *client) *client {
	go func() {
		for request := range slf.todo {
			slf.exec(request)
			idle <- slf
		}
	}()
	return slf
}

// 构建一个客户端
func newClient() *client {
	slf := &client{
		Client: &http.Client{},
		todo: make(chan *http.Request),
	}
	return slf
}