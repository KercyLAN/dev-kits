package file

import (
	"io"
	"os"
)

// 文件读取处理器目标回调函数
type readFileHandlerTargetCallback func(index int, data []byte, err error)

// 文件读取处理器目标结构描述
type readFileHandlerTarget struct {
	master 				*Corps

	index 				int
	file 				*os.File
	start 				int64
	callback			readFileHandlerTargetCallback
}

func (slf *readFileHandlerTarget) Exec() {
	buffer := make([]byte, slf.master.GetConfig().BufferSize)
	if len, err := slf.file.ReadAt(buffer, slf.start); err != nil {
		if err == io.EOF {
			slf.callback(slf.index, buffer[:len], nil)
		}else {
			slf.callback(slf.index, nil, err)
		}
	}else {
		slf.callback(slf.index, buffer[:len], nil)
	}
}

func (slf *readFileHandlerTarget) SetMaster(corps *Corps) Handler {
	slf.master = corps
	return slf
}

// 构建一个文件读取处理器目标实例
func newReadFileHandlerTarget(index int, start int64, file *os.File, callback readFileHandlerTargetCallback) Handler {
	slf := &readFileHandlerTarget{
		index: index,
		start: start,
		file: file,
		callback: callback,
	}
	return slf
}
