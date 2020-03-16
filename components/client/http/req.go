package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// 构建完毕待发出的请求描述
type req struct {
	config 					*conf 						// 配置信息
	bodyBuffer				*bytes.Buffer				// 请求体缓冲器
	bodyWriter 				*multipart.Writer			// 请求体的表单Writer
}

// 决定这个就绪的请求发起同步请求
func (slf *req) Synchronous() (*Result, error) {
	return &Result{}, nil
}


// 决定这个就绪的请求发起异步请求
func (slf *req) Asynchronous() {

}

// 表单中写入一个文件
func (slf *req) CreateFormFile(fieldname string, filePath string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个表单文件Writer
	fileWriter, err := slf.bodyWriter.CreateFormFile(fieldname, file.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return err
	}

}

// 构建一个带发出的请求实例
func newReq(config *conf) *req {
	slf := new(req)
	slf.config = config
	slf.bodyBuffer = &bytes.Buffer{}
	slf.bodyWriter = multipart.NewWriter(slf.bodyBuffer)

	return slf
}