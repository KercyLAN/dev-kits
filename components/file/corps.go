package file

// 文件处理组件结构描述
//
// 文件处理组件file是一个多线程的文件操作组件。
//
// 在操作文件的时候只需要传入不同的handler即可进行不同的文件处理方式。
//
// 如果需要自定义handler只需要实现handler即可。
type Corps struct {
	config 			*conf
	idle			chan *filer
}

// 开始处理
//
// 实现了handler接口的结构体视为一个处理器，将根据不同的处理器进行不同的文件处理。
func (slf *Corps) Do(handler Handler) {
	(<- slf.idle).add(handler.SetMaster(slf))
}

// 获取配置信息
func (slf *Corps) GetConfig() *conf {
	return slf.config
}

// 构建一个文件处理组件
//
// 如果不传入config则采用默认配置，如果传入多个则取第一个。
func New(config ...*conf) *Corps {
	var useConf *conf
	if len(config) > 0 {
		useConf = config[0]
	} else {
		useConf = NewConf()
	}

	slf := &Corps{
		config: useConf,
		idle:   make(chan *filer, useConf.MaxFiler),
	}

	if useConf.DefaultIdle > useConf.MaxFiler {
		useConf.DefaultIdle = useConf.MaxFiler
	}
	for i := 0; i < useConf.DefaultIdle; i++ {
		slf.idle <- newFiler().run(slf.idle)
	}

	return slf
}