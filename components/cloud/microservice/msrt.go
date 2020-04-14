package microservice

import (
	"github.com/KercyLAN/dev-kits/components/cloud"
	"sync"
)

// 微服务注册表结构描述
type RegistryTable struct {
	services					sync.Map					// 微服务集合
}

// 获取特定名称的应用程序
func (slf *RegistryTable) GetService(name string) *Service {
	if app, ok := slf.services.Load(name); ok {
		return app.(*Service)
	}
	return nil
}

// 注册应用程序到微服务注册表中
//
// 当发生这个应用程序已经注册过的情况，合并应用程序已记录的节点信息
func (slf *RegistryTable) Registry(service *Service) {
	// 从已注册的服务中查询
	// 如果不存在则正常注册，否则合并
	existService := slf.GetService(service.Name)
	if existService == nil {
		slf.services.Store(service.Name, service)
	}else {
		service.Nodes.Range(func(key, value interface{}) bool {
			host, node := key.(string), value.(cloud.Node)
			// 节点存储
			// todo：可能会有一种特殊情况，两个相同的节点取谁？待考证
			existService.Nodes.Store(host, node)
			return true
		})
	}
}

// 构建微服务注册表实例
func NewMSRT() *RegistryTable {
	slf := &RegistryTable{

	}
	return slf
}