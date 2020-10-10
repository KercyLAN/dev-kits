package application

var Conf *Configurator;

func init() {
	Conf = &Configurator{
		Micro: &Micro{Host: "", Port: 19100, RegistryServerAddress: ""},
	}
}

// 开发套件全局配置器
//
// 配置器会在程序运行时读取运行目录下的“kapp.*”文件，当找到符合的配置文件会进行序列化到配置中。
//
// 当没有找到配置文件的情况下采取默认配置。
type Configurator struct {
	Micro *Micro `json:"micro"`
}
