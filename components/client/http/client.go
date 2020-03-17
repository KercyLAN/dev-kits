package http

import (
	"io/ioutil"
	"net/http"
)

// 网络请求客户端结构描述
type client struct {
	*http.Client							// 原生HTTP客户端
	todo 			chan *Req				// 待办请求
}

// 执行请求
func (slf *client) exec(req *Req) {
	response, err := slf.Client.Do(req.request)
	if err != nil {
		if req.ansyCallBack != nil {
			req.ansyCallBack(nil, err)
		}
		if req.callBack != nil {
			req.callBack(nil, nil, err)
		}
		return
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if req.ansyCallBack != nil {
			req.ansyCallBack(nil, err)
		}
		if req.callBack != nil {
			req.callBack(nil, nil, err)
		}
		return
	}
	defer response.Body.Close()

	// 处理Cookie
	if req.master.config.UseSetCookie {
		for _, cookie := range response.Cookies() {
			if localCookie, exist := req.master.cookies[cookie.Name]; exist {
				localCookie.Expires = cookie.Expires
				localCookie.Path = cookie.Path
				localCookie.Domain = cookie.Domain
				localCookie.Secure = cookie.Secure
				localCookie.HttpOnly = cookie.HttpOnly
				localCookie.Raw = cookie.Raw
				localCookie.Value = cookie.Value
				localCookie.MaxAge = cookie.MaxAge
				localCookie.RawExpires = cookie.RawExpires
				localCookie.SameSite = cookie.SameSite
				localCookie.Unparsed = cookie.Unparsed
			}else {
				req.master.cookies[cookie.Name] = cookie
			}
		}
	}

	if req.ansyCallBack != nil {
		req.ansyCallBack(newResult(bytes, req.request, response), err)
	}
	if req.callBack != nil {
		req.callBack(bytes, response, err)
	}
}

// 添加待办请求
func (slf *client) add(req *Req) {
	slf.todo <- req
}


// 运行客户端
func (slf *client) run(idle chan <- *client) *client {
	go func() {
		for request := range slf.todo {
			request.format()
			slf.exec(request)
			idle <- slf
		}
	}()
	return slf
}

// 设置transport
func (slf *client) setTransport(transport *http.Transport) *client {
	slf.Transport = transport
	return slf
}

// 构建一个客户端
func newClient() *client {
	slf := &client{
		Client: &http.Client{},
		todo: make(chan *Req),
	}
	return slf
}