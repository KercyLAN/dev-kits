package file

// 文件处理组件处理器行为结构定义
type Handler interface {
	// 执行处理器操作
	Exec()
	// 设置文件处理组件
	SetMaster(corps *Corps) Handler
}

