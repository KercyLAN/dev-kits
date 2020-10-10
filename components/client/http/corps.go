package http

import (
	"crypto/tls"
	"github.com/KercyLAN/dev-kits/components/file"
	"net/http"
)

// 支持并发使用的Http客户端
//
// 支持同步和异步请求
type Corps struct {
	config 					*conf					// 配置信息
	fileCorps 				*file.Corps				// 文件处理组件
	idle 					chan *client			// 空闲的客户端管道
	headers 				map[string]string		// 永久Header
	cookies 				map[string]*http.Cookie // 永久Cookie
	transport 				*http.Transport
}

// 构建请求
func (slf *Corps) Do(method Method, url string) *Req {
	return newReq(slf, method, url)
}

// 构建一个Http客户端
//
// 允许传入配置信息，当传入多个的情况下仅接受第一个
func New(config ...*conf) *Corps {
	// 配置初始化
	var useConf *conf
	if len(config) > 0 {
		useConf = config[0]
	} else {
		useConf = NewConf()
	}

	fileCorpsConfig := file.NewConf()
	fileCorpsConfig.BufferSize = useConf.BufferSize
	fileCorpsConfig.MaxFiler = useConf.MaxFiler
	fileCorpsConfig.DefaultIdle = useConf.DefaultIdle

	// 实例初始化
	slf := &Corps{
		config: useConf,
		idle: make(chan *client, useConf.MaxClient),
		fileCorps: file.New(fileCorpsConfig),
		headers: map[string]string{},
		cookies: map[string]*http.Cookie{},
		transport: &http.Transport{},
	}
	slf.transport.TLSClientConfig = &tls.Config{}

	// 根据配置构建空闲客户端
	for i := 0; i< useConf.DefaultIdle; i++ {
		slf.idle <- newClient().setTransport(slf.transport).run(slf.idle)
	}

	return slf
}