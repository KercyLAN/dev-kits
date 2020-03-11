// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-3 0:20

package upload

import (
	"mime/multipart"
)

const errMultipleValue = "files and hash values must be one-to-one"

// 对请求中FileGroup的描述。
type FileGroup map[string][]*multipart.FileHeader

// 对请求中ValueGroup的描述。
type ValueGroup map[string][]string

// 返回groupName组的成员数量。
func (slf FileGroup) GroupLen(groupName string) int {
	return len(slf[groupName])
}

// 返回groupName组的成员数量。
func (slf ValueGroup) GroupLen(groupName string) int {
	return len(slf[groupName])
}

// 遍历所有的FileHeader反馈到hook中
//
// 当hook返回的err为value存在多个的时候的时候，整个遍历将终止。
//
// 当hook返回的err为其他如文件创建失败等错误的时候，会进入下一个循环，并且将信息传入failHook中。
//
// 当failHook返回true，则表示直接终止遍历。
func (slf FileGroup) EachAll(
	hook func(groupName string, groupIndex int, fileHeader *multipart.FileHeader) error,
	failHook func(groupName string, groupIndex int, fileName string, fileSize int64, err error) bool) error {
	for groupName, fileHeaders := range slf {
		for i, fileHeader := range fileHeaders {
			err := hook(groupName, i, fileHeader);
			if err != nil {
				if err.Error() == errMultipleValue {
					return err
				}else {
					if failHook(groupName, i, fileHeader.Filename, fileHeader.Size, err) {
						return err
					}
				}
			}
		}
	}
	return nil
}

// 检查groupName这个组是否存在。
func (slf FileGroup) Exist(groupName string) bool {
	_, ok := slf[groupName]
	return ok
}

// 检查groupName这个组是否存在。
func (slf ValueGroup) Exist(groupName string) bool {
	_, ok := slf[groupName]
	return ok
}

// 遍历所有groupName组下的成员反馈到hook中。
func (slf FileGroup) Each(groupName string, hook func(i int, fileHeader *multipart.FileHeader)) {
	for i, fileHeader := range slf[groupName]{
		hook(i, fileHeader)
	}
}

// 遍历所有groupName组下的成员反馈到hook中。
func (slf ValueGroup) Each(groupName string, hook func(i int, value string)) {
	for i, value := range slf[groupName]{
		hook(i, value)
	}
}