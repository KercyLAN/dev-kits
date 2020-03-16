package http


// 支持并发使用的Http客户端
//
// 支持同步和异步请求
type Corps struct {
	config 			*conf					// 配置信息
	idle 			chan *client			// 空闲的客户端管道
}

// 构建请求
func (slf *Corps) Do() *req {
	return nil
}

// 构建一个Http客户端
//
// 允许传入配置信息，当传入多个的情况下仅接受第一个
func New(config ...*conf) *Corps {
	var useConf *conf
	if len(config) > 0 {
		useConf = config[0]
	} else {
		useConf = NewConf()
	}

	slf := &Corps{
		config: useConf,
		idle: make(chan *client, useConf.MaxClient),
	}

	// 根据配置构建空闲客户端
	for i := 0; i< useConf.DefaultIdle; i++ {
		slf.idle <- newClient().run(slf.idle)
	}

	return slf
}