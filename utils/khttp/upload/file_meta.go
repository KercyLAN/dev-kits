// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 23:00

/**
包upload包含了对http文件上传的封装
 */
package upload

import (
	"github.com/KercyLAN/dev-kits/application"
	"github.com/KercyLAN/dev-kits/utils/khttp/cloud"
	uuid "github.com/satori/go.uuid"
	"path"
	"time"
)

// HTTP上传的文件的文件元信息。
type FileMeta struct {
	id 						string					// 文件ID(uuid)
	name 					string					// 文件名称
	storageName 			string					// 在文件系统中的名称
	size 					int64					// 文件大小
	ext 					string					// 文件扩展名
	uploadAt 				time.Time				// 文件上传时间
	hash					string					// 文件的hash
	node 					cloud.Node				// 文件所在节点
	isFastUpload			bool					// 是否为快速上传返回的文件元信息
}

// 检查这个文件是否是快速上传的。
func (slf *FileMeta) IsFastUpload() bool {
	return slf.isFastUpload
}

// 获取文件id。
func (slf *FileMeta) GetId() string {
	return slf.id
}

// 获取文件的真实名称。
func (slf *FileMeta) GetName() string {
	return slf.name
}

// 获取该文件在文件系统中的存储的文件名。
func (slf *FileMeta) GetStorageName() string {
	return slf.storageName
}

// 获取文件占用的空间大小。
func (slf *FileMeta) GetSize() int64 {
	return slf.size
}

// 获取文件的拓展名。
func (slf *FileMeta) GetExt() string {
	return slf.ext
}

// 获取文件的上传时间。
func (slf *FileMeta) GetUploadAt() time.Time {
	return slf.uploadAt
}

// 获取文件的hash值。
func (slf *FileMeta) GetHash() string {
	return slf.hash
}

// 获取文件所对应的节点信息。
func (slf *FileMeta) GetNode() cloud.Node {
	return slf.node
}

// 构建一个完备的文件元信息。
func NewFileMeta(name string, size int64) *FileMeta {
	uid := uuid.NewV4().String()
	ext := path.Ext(name)
	this := &FileMeta{
		id: uid,
		name: name,
		storageName: uid + ext,
		size: size,
		ext: ext,
		uploadAt: time.Now(),
		node: application.Get().Node,
	}
	return this
}

