// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-3 11:51

package upload

import (
	"errors"
	"github.com/KercyLAN/dev-kits/utils/kstreams"
	"os"
	"strconv"
)

// 分块上传处理器
//
// 对FileUploadReceiver的拓展，提供了对分块上传的需求解决。
type partProcessor struct {
	task 					map[string]map[int][]byte // 分块上传任务
	taskTemp				map[string]map[int]string // 分块上传任务（文件存储）
	taskPartCounter 		map[string]int             // 任务块计数器，统计一个任务总共有多少个块
	taskFast				map[string]*FileMeta      // 任务快速上传记录
}


// 检查指定hash的任务是否已经完成。
func (slf *partProcessor) IsFinish(hash string) bool {
	if len(slf.task[hash]) == slf.taskPartCounter[hash] && slf.taskPartCounter[hash] != 0 {
		return true
	}
	return false
}

// 创建一个分块上传任务
//
// 由于分块上传是由多个文件组成一个文件，所以hash表示了他们共同指向的文件hash。
//
// 具体的分块数量由partNumber决定。
func (slf *partProcessor) NewPartTask(hash string, partNumber int) {
	slf.task[hash] = map[int][]byte{}
	slf.taskTemp[hash] = map[int]string{}
	slf.taskPartCounter[hash] = partNumber
	slf.taskFast[hash] = nil
}

// 返回特定hash的分块上传任务已上传的分块序号。
func (slf *partProcessor) GetPartUploadSuccess(hash string) []int {
	parts := make([]int, 0)
	for partSort := range slf.task[hash] {
		parts = append(parts, partSort)
	}
	return parts
}

// 清理指定任务组数据。
func (slf *partProcessor) clear(hash string) {
	delete(slf.task, hash)
	delete(slf.taskPartCounter, hash)
	delete(slf.taskFast, hash)
	for _, path := range slf.taskTemp[hash] {
		if path != "" {
			os.Remove(path)
		}
	}
	delete(slf.taskTemp, hash)

}

// 添加已经下载的块到指定hash的任务中。
func (slf *partProcessor) addPartPath(hash string, partIndex string, part string) error {
	index, err := strconv.ParseInt(partIndex, 10, 64)
	if err != nil {
		return err
	}
	slf.task	[hash][int(index)] = make([]byte, 0)
	slf.taskTemp[hash][int(index)] = part
	return nil
}

// 添加已经下载的块到指定hash的任务中。
func (slf *partProcessor) addPart(hash string, partIndex string, part []byte, receiver *FileUploadReceiver) error {
	if slf.getPartSize(hash) + int64(len(part)) > receiver.partFileMaxSize {
		slf.clear(hash)
		return errors.New("file size out range of")
	}

	index, err := strconv.ParseInt(partIndex, 10, 64)
	if err != nil {
		return err
	}
	slf.task	[hash][int(index)] = part
	slf.taskTemp[hash][int(index)] = ""

	return nil
}

// 返回指定任务组文件总大小。
func (slf *partProcessor) getPartSize(hash string) int64 {
	size := int64(0)
	kstreams.EachMapSort(slf.task[hash], func(partIndex int, data []byte) {
		size += int64(len(data))
	})
	return size
}