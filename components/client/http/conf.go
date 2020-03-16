package http

// 客户端配置描述
type conf struct {
	MaxClient 			int				// 最大客户端数量
	DefaultIdle			int 			// 默认空闲客户端数量
	FileBufferSize		int 			// 文件分块读取每块大小
	FileReaderCount		int 			// 最多支持多少个线程共同读取一个文件
}

// 构建一个配置实例
func NewConf() *conf {
	slf := &conf{
		MaxClient:   50,
		DefaultIdle: 10,
		FileBufferSize: 10 * 1024 * 1024,
		FileReaderCount: 10,
	}
	return slf
}