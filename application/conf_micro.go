package application

import "fmt"

// 微服务相关配置
type Micro struct{
	cache 						map[string]string

	Host 						string `json:"host"`						// 监听的Host
	Port 						int `json:"port"`							// 监听的端口
	RegistryServerAddress		string `json:"registry_server_address"`		// 注册中心地址
}

func (slf *Micro) Address() string {
	if slf.cache == nil {
		slf.cache = map[string]string{}
		address := fmt.Sprint(slf.Host, ":", slf.Port)
		slf.cache["address"] = address
		return address
	}
	if val, ok := slf.cache["address"]; ok {
		return val
	}else {
		address := fmt.Sprint(slf.Host, ":", slf.Port)
		slf.cache["address"] = address
		return address
	}
}
