// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-1 12:15

/**
包cloud包含了对云端程序的通用描述及功能集合
 */
package cloud

// 计算机节点的通用描述接口
type Node interface {
	// 获取节点id，
	GetId() string
	// 获取节点host。
	GetHost() string
}