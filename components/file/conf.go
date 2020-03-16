package file

// 文件操作组件配置描述
type conf struct {
	MaxFiler 			int				// 最大Filer数量
	DefaultIdle			int 			// 默认空闲Filer数量
	BufferSize			int64 			// 缓冲区大小
}

// 构建一个文件操作组件配置信息实例
func NewConf() *conf {
	slf := &conf{
		MaxFiler: 10,
		DefaultIdle: 5,
		BufferSize: 10 * 1024 * 1024,
	}
	return slf
}