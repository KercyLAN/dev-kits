// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 21:56

package kproperties

// PropertiesInterface 表示了一个完整的Properties需要实现的函数
type PropertiesInterface interface {
	// GetString 返回这个properties中key对应的value转换为string类型后的值
	GetString(key string) string
	// GetInt 返回这个properties中key对应的value转换为int类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetInt(key string) (int, error)
	// GetInt8 返回这个properties中key对应的value转换为int8类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetInt8(key string) (int8, error)
	// GetInt16 返回这个properties中key对应的value转换为int16类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetInt16(key string) (int16, error)
	// GetInt32 返回这个properties中key对应的value转换为int32类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetInt32(key string) (int32, error)
	// GetInt64 返回这个properties中key对应的value转换为int64类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetInt64(key string) (int64, error)
	// GetUint 返回这个properties中key对应的value转换为uint类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetUint(key string) (uint, error)
	// GetUint8 返回这个properties中key对应的value转换为uint8类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetUint8(key string) (uint8, error)
	// GetUint16 返回这个properties中key对应的value转换为uint16类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetUint16(key string) (uint16, error)
	// GetUint32 返回这个properties中key对应的value转换为uint32类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetUint32(key string) (uint32, error)
	// GetUint64 返回这个properties中key对应的value转换为uint64类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetUint64(key string) (uint64, error)
	// GetBool 返回这个properties中key对应的value转换为bool类型后的值
	//
	// 默认实现中"1", "true", "ok", "yes", "sure", "affirm", "has", "success"均为true，反之为false
	GetBool(key string) bool
	// GetInterface 返回这个properties中key对应的value转换为interface类型后的值
	GetInterface(key string) interface{}
	// GetFloat32 返回这个properties中key对应的value转换为float32类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789.123”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetFloat32(key string) (float32, error)
	// GetFloat64 返回这个properties中key对应的value转换为float64类型后的值
	//
	// 默认的实现中为了方便查看，允许在value中填写“123,456,789.123”这种格式。
	// 默认的实现中会自动将value中的空格和“,”分隔符全部排除。
	GetFloat64(key string) (float64, error)
	// Each 遍历所有的kv对
	Each(eachFunc func(key string, value string))
	// HasKey 返回是否存在Key
	HasKey(key string) bool
}
