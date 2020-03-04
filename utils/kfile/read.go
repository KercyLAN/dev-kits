// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 20:50

/**
包kfile包含了一些针对文件处理的封装
 */
package kfile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// 单次读取文件
//
// 一次性对整个文件进行读取，小文件读取可以很方便的一次性将文件内容读取出来，而大文件读取会造成性能影响。
func ReadOnce(filePath string) ([]byte, error) {
	if file, err := os.Open(filePath); err != nil {
		return nil, err
	}else {
		defer file.Close()
		return ioutil.ReadAll(file)
	}
}

// 分块读取文件
//
// 将filePath路径对应的文件数据并将读到的每一部分传入hook函数中，当过程中如果产生错误则会返回error。
//
// 分块读取可以在读取速度和内存消耗之间有一个很好的平衡。
//
// 具体使用方法如下：
//	func main() {
//		byteList := make([][]byte, 0)
//		err := ReadBlockHook("filepath...", 1024, func(data []byte) {
//			byteList = append(byteList, data)
//		})
//		if err == nil {
//			fileData := bytes.Join(byteList, make([]byte, 0))
//		}
//		...
//	}
func ReadBlockHook(filePath string, bufferSize int, hook func(data []byte)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := make([]byte, bufferSize)
	bufferReader := bufio.NewReader(file)
	for {
		successReadSize, err := bufferReader.Read(buffer)
		hook(buffer[:successReadSize])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// 逐行读取文件
//
// 将filePath路径对应的文件数据逐行读取，并将每一行读取到的数据传入hook中，过程中如果发生错误则会返回error。
//
// 逐行读取的性能会稍慢，仅占用极少的内存空间，同时由于性能慢的原因也极其不推荐用于读取非文本类文件。
//
// 具体使用方法如下：
//	func main (){
//		err := kfile.ReadLineHook("filepath...", func(data []byte) {
//			fmt.Println(string(data))
//		})
//		...
//	}
func ReadLineHook(filePath string, hook func(data []byte)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	bufferReader := bufio.NewReader(file)
	for {
		line, err := bufferReader.ReadBytes('\n')
		hook(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}