package file

import (
	"bytes"
	"math"
	"os"
	"sync"
)

// 读取文件处理器回调函数
type ReadFileHandlerCallBack func(data []byte, err error)

// 读取文件处理器结构描述
//
// 该处理器会对文件进行分块读取，在最终读取完毕后进行汇总。
//
// 当读取的文件大小超过系统剩余内存会发生不可预知的情况
// todo：超出系统内存的文件读取待优化
type ReadFileHandler struct {
	master 				*Corps						// 文件处理组件

	filePath 			string						// 需要操作的文件路径
	callback 			ReadFileHandlerCallBack		// 回调函数
	blocks 				[][]byte					// 所有文件块
	successCount		int							// 读取成功的文件快数量
	callErr 			error						// 嵌套的处理器错误信息
}


func (slf *ReadFileHandler) Exec() {
	file, err := os.Open(slf.filePath)
	if err != nil {
		slf.callback(nil, err)
	}
	defer file.Close()

	fileInfo, err := os.Stat(slf.filePath)
	if err != nil {
		slf.callback(nil, err)
	}
	slf.blocks = make([][]byte, int(math.Ceil(float64(fileInfo.Size()) / float64(slf.master.GetConfig().BufferSize))))

	var wait sync.WaitGroup
	for i := int64(0); i < int64(len(slf.blocks)); i++ {
		wait.Add(1)
		start := i * slf.master.GetConfig().BufferSize
		slf.master.Do(newReadFileHandlerTarget(int(i), start, file, func(index int, data []byte, err error){
				if err != nil {
					slf.callErr = err
					return
				}
				slf.blocks[index] = data
				slf.successCount++

				if slf.successCount == len(slf.blocks) {
					slf.callback(bytes.Join(slf.blocks, make([]byte, 0)), nil)
				}

				wait.Done()
		}))
	}

	wait.Wait()
}

func (slf *ReadFileHandler) SetMaster(corps *Corps) Handler {
	slf.master = corps
	return slf
}

// 构建一个文件读取处理器实例
func NewReadFileHandler(filePath string, callback ReadFileHandlerCallBack) Handler {
	slf := &ReadFileHandler{
		filePath: filePath,
		callback: callback,
	}
	return slf
}