package microservice

import (
	"github.com/KercyLAN/dev-kits/components/cloud"
	"sync"
)

// 微服务程序结构描述
//
// 应用程序包含了多个实例节点信息
type Service struct {
	Name 					string								// 应用程序名称
	Nodes					sync.Map							// 这个应用程序包含的节点(使用host作为key避免重复)
}

// 添加节点
func (slf *Service) Add(node cloud.Node) {
	slf.Nodes.Store(node.GetHost(), node)
}

// 构建一个微服务应用程序
func NewApplication(name string, nodes ...cloud.Node) *Service {
	slf := &Service{
		Name: name,
	}
	return slf
}