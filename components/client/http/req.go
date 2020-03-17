package http

import (
	"bytes"
	"encoding/json"
	"github.com/KercyLAN/dev-kits/components/file"
	"github.com/KercyLAN/dev-kits/utils/kstr"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

// 回调函数
type reqCallback func(data []byte, response *http.Response, err error)

// 异步回调
type CallBack func(result *Result, err error)

// 构建完毕待发出的请求描述
type req struct {
	request 				*http.Request					// Request
	master 					*Corps 							// http客户端组件
	bodyBuffer				*bytes.Buffer					// 请求体缓冲器
	bodyWriter 				*multipart.Writer				// 请求体的表单Writer
	fileCount 				int								// 请求中的文件数量
	contentType				string							// 内容类型
	contentTypeParam		[]string						// Content-Type参数，如charset=UTF-8
	json 					[]byte 							// json
	content					string							// text
	params 					map[string][]string				// 携带的请求参数

	callBack 				reqCallback						// 回调函数
	ansyCallBack 			CallBack 						// 异步回调函数
	waitGroup 				sync.WaitGroup
	Err 					error							// 请求是否存在异常
}

// 决定这个就绪的请求发起同步请求
func (slf *req) SynchronousExec() (*Result, error) {
	if slf.Err != nil {
		return nil, slf.Err
	}

	var result *Result
	var callErr error
	slf.waitGroup.Add(1)
	slf.callBack = func(data []byte, response *http.Response, err error) {
		if err != nil {
			callErr = err
		} else {
			result = newResult(data, slf.request, response)
		}
		slf.waitGroup.Done()
	}

	(<- slf.master.idle).add(slf)
	slf.waitGroup.Wait()
	return result, callErr
}


// 决定这个就绪的请求发起异步请求
func (slf *req) AsynchronousExec(callback CallBack) {
	if slf.Err != nil {
		callback(nil, slf.Err)
		return
	}

	slf.ansyCallBack = callback
	(<- slf.master.idle).add(slf)
}

// 设置请求Cookie
func (slf *req) SetCookie(cookie *http.Cookie) *req {
	slf.master.cookies[cookie.Name] = cookie
	return slf
}

// 设置请求永久Header
//
// 这样设置的Header将会在整个Corps实例中均存在
func (slf *req) SetHeaderForever(name, value string) *req {
	slf.master.headers[name] = value
	return slf
}

// 设置请求临时Header
func (slf *req) SetHeader(name, value string) *req {
	if slf.request != nil {
		if name == "Content-Type" {
			return slf
		}
		slf.request.Header.Set(name, value)
	}
	return slf
}

// 设置请求发送JSON数据
func (slf *req) Json(json []byte) {
	slf.contentType = "application/json"
	slf.json = json
}

// 设置请求发送特定类型的文本信息
func (slf *req) Content(content string, contentType string) {
	slf.contentType = contentType
	slf.content = content
}

// 设置请求发送JSON数据
func (slf *req) JsonEntity(jsonEntity interface{}) *req {
	json, err := json.Marshal(jsonEntity)
	if err != nil {
		slf.Err = err
	}else {
		slf.json = json
		slf.contentType = "application/json"
	}
	return slf
}

// 设置请求Content-Type
func (slf *req) SetContentType(contentType string) *req {
	slf.contentType = contentType
	return slf
}

// 添加请求Content-Type参数
func (slf *req) AddContentTypeParam(param string) *req {
	slf.contentTypeParam = append(slf.contentTypeParam, param)
	return slf
}

// 格式化请求
func (slf *req) format() {
	if slf.request == nil {
		return
	}

	// 生成Content-Type参数
	var contentTypeParamBunch string
	for _, param := range slf.contentTypeParam {
		contentTypeParamBunch += "; " + param
	}

	// 格式化header
	for name, value := range slf.master.headers {
		slf.request.Header.Set(name, value)
	}

	// 格式化Cookie
	for _, cookie := range slf.master.cookies {
		if existCookie, err := slf.request.Cookie(cookie.Name); err != nil {
			slf.request.AddCookie(cookie)
		}else {
			existCookie.Expires = cookie.Expires
			existCookie.Path = cookie.Path
			existCookie.Domain = cookie.Domain
			existCookie.Secure = cookie.Secure
			existCookie.HttpOnly = cookie.HttpOnly
			existCookie.Raw = cookie.Raw
			existCookie.Value = cookie.Value
			existCookie.MaxAge = cookie.MaxAge
			existCookie.RawExpires = cookie.RawExpires
			existCookie.SameSite = cookie.SameSite
			existCookie.Unparsed = cookie.Unparsed
		}
	}

	if strings.ToUpper(slf.request.Method) != "GET" {
		if slf.fileCount > 0 {
			// 文件请求
			if slf.Err != nil {
				return
			}
			slf.request.Header.Set("Content-Type", slf.bodyWriter.FormDataContentType() + contentTypeParamBunch)
			for key, values := range slf.params {
				for _, value := range values {
					slf.Err = slf.bodyWriter.WriteField(key, value)
					if slf.Err != nil {
						return
					}
				}
			}
			slf.request.Body = ioutil.NopCloser(slf.bodyBuffer)
			slf.request.ContentLength = int64(slf.bodyBuffer.Len())
		}else {
			// 其他请求
			slf.request.Header.Set("Content-Type",  slf.contentType + contentTypeParamBunch)
			switch slf.contentType {
			case "application/json":
				setReader(slf.request, bytes.NewBuffer(slf.json))
			case "text/plain", "application/javascript", "application/xml", "text/xml", "text/html":
				setReader(slf.request, strings.NewReader(slf.content))
			default:
				slf.request.Header.Set("Content-Type", "application/x-www-form-urlencoded" + contentTypeParamBunch)
				params := &url.Values{}
				for k, vs := range slf.params {
					for _, v := range vs {
						params.Add(k, v)
					}
				}
				setReader(slf.request, strings.NewReader(params.Encode()))
			}
		}

		if slf.request.GetBody != nil && slf.request.ContentLength == 0 {
			slf.request.Body = http.NoBody
			slf.request.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		}
	}
}

// 设置client的reader
func setReader(request *http.Request, reader io.Reader)  {
	rc, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		rc = ioutil.NopCloser(reader)
	}
	request.Body = rc

	switch v := reader.(type) {
	case *bytes.Buffer:
		request.ContentLength = int64(v.Len())
		buf := v.Bytes()
		request.GetBody = func() (io.ReadCloser, error) {
			r := bytes.NewReader(buf)
			return ioutil.NopCloser(r), nil
		}
	case *bytes.Reader:
		request.ContentLength = int64(v.Len())
		snapshot := *v
		request.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return ioutil.NopCloser(&r), nil
		}
	case *strings.Reader:
		request.ContentLength = int64(v.Len())
		snapshot := *v
		request.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return ioutil.NopCloser(&r), nil
		}
	}
	if request.GetBody != nil && request.ContentLength == 0 {
		request.Body = http.NoBody
		request.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
	}
}

// 请求URL中添加参数
func (slf *req) AddUrlParam(key, value string) *req {
	if slf.request != nil {
		v := slf.request.URL.Query()
		v.Add(key, value)
		slf.request.URL.RawQuery = v.Encode()
	}
	return slf
}

// 请求URL中添加参数串
func (slf *req) AddUrlParamBunch(bunch string) *req {
	if slf.request != nil {
		v := slf.request.URL.Query()
		for _, slice := range strings.Split(bunch, "&") {
			v.Add(kstr.KV(slice, "="))
		}
		slf.request.URL.RawQuery = v.Encode()
	}
	return slf
}

// 表单中添加参数
func (slf *req) AddFormParam(key, value string) *req {
	if slf.params[key] == nil {
		slf.params[key] = make([]string, 0)
	}
	slf.params[key] = append(slf.params[key], value)
	return slf
}

// 表单中添加参数串
func (slf *req) AddFormParamBunch(bunch string) *req {
	for _, slice := range strings.Split(bunch, "&") {
		k, v := kstr.KV(slice, "=")
		if slf.params[k] == nil {
			slf.params[k] = make([]string, 0)
		}
		slf.params[k] = append(slf.params[k], v)
	}
	return slf
}

// 表单中写入一个文件
func (slf *req) CreateFormFile(fieldName string, filePath string) *req {
	// 打开文件
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		slf.Err = err
		return slf
	}

	// 创建一个表单文件Writer
	fileWriter, err := slf.bodyWriter.CreateFormFile(fieldName, fileInfo.Name())
	if err != nil {
		slf.Err = err
		return slf
	}

	// 使用文件处理组件读取文件，并进行阻塞，不进行异步处理
	var fileCorpsError error
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	slf.master.fileCorps.Do(file.NewReadFileHandler(filePath, func(data []byte, err error) {
		if err != nil {
			fileCorpsError = err
			return
		}
		fileWriter.Write(data)
		waitGroup.Done()
	}))
	waitGroup.Wait()

	if fileCorpsError != nil {
		slf.Err = fileCorpsError
		return slf
	}
	slf.fileCount++
	return slf
}

// 构建一个带发出的请求实例
func newReq(master *Corps, method method, url string) *req {
	slf := new(req)
	slf.master = master
	slf.bodyBuffer = &bytes.Buffer{}
	slf.bodyWriter = multipart.NewWriter(slf.bodyBuffer)
	slf.params = make(map[string][]string)
	slf.contentTypeParam = make([]string, 0)

	request, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		slf.Err = err
	}else {
		slf.request = request
	}
	return slf
}