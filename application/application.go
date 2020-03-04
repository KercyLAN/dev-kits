// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-1 12:22

/**
包application对应用程序进行了通用的描述以及提供了一些全局性的功能集。
 */
package application

import "github.com/KercyLAN/dev-kits/utils/khttp/cloud"

// 全局唯一的应用程序实例
var instance *application

// 应用程序通用描述功能集。
type application struct {
	Node 				cloud.Node				// 该应用程序的节点信息
}

// 初始化应用程序
//
// 初始化来自kercylan-lib提供的应用程序通用描述功能集。
//
// 由于application是可能会被整个kercylan-lib的各个实现所使用，且kercylan-lib是一个支持云端应用构建的支持库，
// 所以需要提供一个节点信息来描述运行这个应用程序的计算机在云端所表示的节点信息。
//
// 我们强烈建议在main包的init()函数下对其进行初始化，具体如下：
//	func init() {
//		// 硬编码Host
//		application.Initialize(cloud.NewNode("127.0.0.1"))
//		// 通过knet.IP()函数获取到本机IP地址进行赋值
//		application.Initialize(cloud.NewNode(knet.IP().String()))
//	}
//
// 当然，关于node的来源如果有更好的办法，只需要编写一个实现cloud.Node接口的结构体即可。
func Initialize(node cloud.Node) {
	instance = &application{
		Node: node,
	}
}

// 获取应用程序实例，如果尚未初始化则会返回产生panic。
func Get() *application {
	if instance == nil {
		panic("the application is not yet initialized, use \"application.Initialize ()\"")
	}
	return instance
}