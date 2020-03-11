// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-1 12:11

package cloud

import (
	uuid "github.com/satori/go.uuid"
)

// 计算机节点的kercylan-lib实现。
type KNode struct {
	Id 						string					// 节点ID（uuid）
	Host 					string					// 节点Host
}

// 构建一个KNode实例
//
// 会以host来创建一个KNode，这个节点描述了这个应用程序在云端的一切信息。
//
// 由于KNode在创建时会自己产生一个ID，所以极其建议将其使用数据库或其他方式进行存储。
//
// 这样在各种分布式、集群场景中可以确保通过数据库使用id检索到节点信息，也可以使用节点信息检索到节点id。
func NewNode(host string) *KNode {
	this := &KNode{
		Id: uuid.NewV4().String(),
		Host:host,
	}
	return this
}

func (slf *KNode) GetId() string {
	return slf.Id
}

func (slf *KNode) GetHost() string {
	return slf.Host
}
