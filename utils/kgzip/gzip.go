package kgzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// 对数据进行GZip压缩，返回bytes.Buffer和错误信息
func GZipCompress(data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	// 写入压缩
	_, err := gzipWriter.Write(data)
	if err != nil {
		return buf, err
	}
	// 关闭压缩
	if err := gzipWriter.Close(); err != nil {
		return buf, err
	}
	return buf, nil
}

// 对已进行GZip压缩的数据进行解压缩，返回字节数组及错误信息
func UnGZipCompress(dataByte []byte) ([]byte, error) {
	data := *bytes.NewBuffer(dataByte)
	// 读取压缩过的数据
	gzipReader, err := gzip.NewReader(&data)
	if err != nil {
		return nil, err
	}
	// 读取数据
	result, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	// 关闭读取
	if err := gzipReader.Close(); err != nil {
		return nil, err
	}

	return result, nil
}
