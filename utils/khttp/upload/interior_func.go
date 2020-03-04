// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-1 12:04

package upload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 保存文件到本地返回FileMeta。
func saveFile(receiver *FileUploadReceiver, fileHeader *multipart.FileHeader, storageDir string, bufferSize int64) (*FileMeta, string, error) {
	fileMeta := NewFileMeta(fileHeader.Filename, fileHeader.Size)
	var uploadFileHash string
	file, err := os.Create(fmt.Sprintf("%v\\%v%v", storageDir, fileMeta.id, fileMeta.ext))
	if err != nil {
		return nil, uploadFileHash, err
	}
	defer file.Close()

	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, uploadFileHash, err
	}

	fileBytes := make([][]byte, 0)
	nowLen := int64(0)
	for {
		buffer := make([]byte, bufferSize)
		readLen, readErr := formFile.ReadAt(buffer, nowLen)
		nowLen += int64(readLen)

		if receiver.completenessCheckHandler != nil {
			fileBytes = append(fileBytes, buffer[0:readLen])
		}
		if _, err = file.Write(buffer[0:readLen]); err != nil {
			os.Remove(fmt.Sprintf("%v\\%v%v", storageDir, fileMeta.id, fileMeta.ext))
			return nil, uploadFileHash, err
		}

		if readErr != nil {
			if readErr == io.EOF {
				break
			}else {
				return nil, uploadFileHash, readErr
			}
		}

	}

	if receiver.completenessCheckHandler != nil {
		uploadFileHash = receiver.completenessCheckHandler(bytes.Join(fileBytes, make([]byte, 0)))
	}

	return fileMeta, uploadFileHash, nil
}

// 返回form中的file组和value组。
func readForm(reader *multipart.Reader, maxMemory int64) (FileGroup, ValueGroup, error) {
	form, err := reader.ReadForm(maxMemory)
	if err != nil {
		return nil, nil, err
	}
	return form.File, form.Value, err
}

// 返回符合分块上传的form中的信息。
func readFormPart(reader *multipart.Reader, maxMemory int64) (*multipart.FileHeader, string, int64, string, error) {
	form, err := reader.ReadForm(maxMemory)
	var fileHeader *multipart.FileHeader
	var hash string
	var partIndex string
	var fileSize int64
	if err != nil {
		return fileHeader, hash, fileSize, partIndex, err
	}
	for _, fileHeaders := range form.File {
		for _, header := range fileHeaders {
			fileHeader = header
		}
	}
	for _, values := range form.Value {
		for _, value := range values {
			valuePart := strings.SplitN(value, ";", 2)
			if len(valuePart) != 3 {
				return fileHeader, hash, fileSize, partIndex, errors.New("incorrect text parameter, should be formatted as \"fileHash;PartIndex;fileSize\"")
			}
			hash = valuePart[0]
			partIndex = valuePart[1]
			if size, err := strconv.ParseInt(valuePart[2], 10, 64); err != nil {
				return fileHeader, hash, fileSize, partIndex, errors.New("a file size flag for an exception")
			}else {
				fileSize = size
			}
		}
	}
	if fileHeader == nil {
		return fileHeader, hash, fileSize, partIndex, errors.New("no files were found")
	}
	if hash == "" || partIndex == "" {
		return fileHeader, hash, fileSize, partIndex, errors.New("incorrect text parameter, should be formatted as \"fileHash;PartIndex:fileSize\"")
	}
	return fileHeader, hash, fileSize, partIndex, err
}

// 返回FileUploadReceiver所需的内容
//
// FileUploadReceiver仅支持来自POST请求的调用，如果不是，则会返回error。
func initRequest(request *http.Request) ([]*FileMeta, *multipart.Reader, error) {
	if request.Method != "POST" {
		return nil, nil, errors.New("FileUploadReceiver handler upload request. method must is POST")
	}
	multipartReader, err := request.MultipartReader(); if err != nil {
		return nil, nil, err
	}
	return make([]*FileMeta, 0), multipartReader, nil
}

// 返回FileUploadReceiver分块上传所需的内容
//
// FileUploadReceiver仅支持来自POST请求的调用，如果不是，则会返回error。
func initRequestPart(request *http.Request) (*multipart.Reader, error) {
	if request.Method != "POST" {
		return nil, errors.New("FileUploadReceiver handler upload request. method must is POST")
	}
	multipartReader, err := request.MultipartReader(); if err != nil {
		return nil, err
	}
	return multipartReader, nil
}