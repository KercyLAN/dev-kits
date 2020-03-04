// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 20:45

package kproperties

import (
	"errors"
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kfile"
	"github.com/KercyLAN/dev-kits/utils/kobj"
	"github.com/KercyLAN/dev-kits/utils/kstr"
	"github.com/KercyLAN/dev-kits/utils/kstreams"
	"strings"
)

// Properties 表示了对properties格式文件的数据结构体
type Properties struct {
	details 		map[string]string			// 细节
}

func (slf *Properties) HasKey(key string) bool {
	_, ok := slf.details[key]
	return ok
}

func (slf *Properties) Each(eachFunc func(key string, value string)) {
	for k, v := range slf.details{
		eachFunc(k, v)
	}
}

func (slf *Properties) GetString(key string) string {
	return slf.details[key]
}

func (slf *Properties) GetInt(key string) (int, error) {
	return kstr.FormatSpeedyInt(slf.details[key])
}

func (slf *Properties) GetInt8(key string) (int8, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return int8(value), err
}

func (slf *Properties) GetInt16(key string) (int16, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return int16(value), err
}

func (slf *Properties) GetInt32(key string) (int32, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return int32(value), err
}

func (slf *Properties) GetInt64(key string) (int64, error) {
	return kstr.FormatSpeedyInt64(slf.details[key])
}

func (slf *Properties) GetUint(key string) (uint, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return uint(value), err
}

func (slf *Properties) GetUint8(key string) (uint8, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return uint8(value), err
}

func (slf *Properties) GetUint16(key string) (uint16, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return uint16(value), err
}

func (slf *Properties) GetUint32(key string) (uint32, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return uint32(value), err
}

func (slf *Properties) GetUint64(key string) (uint64, error) {
	value, err := kstr.FormatSpeedyInt(slf.details[key])
	return uint64(value), err
}

func (slf *Properties) GetBool(key string) bool {
	switch slf.details[key] {
	case "1", "true", "ok", "yes", "sure", "affirm", "has", "success":
		return true
	default:
		return false
	}
}

func (slf *Properties) GetInterface(key string) interface{} {
	return slf.details[key]
}

func (slf *Properties) GetFloat32(key string) (float32, error) {
	value, err := kstr.FormatSpeedyFloat32(slf.details[key])
	return float32(value), err
}

func (slf *Properties) GetFloat64(key string) (float64, error) {
	return kstr.FormatSpeedyFloat64(slf.details[key])
}

// New 返回由propertiesFilePath路径的properties文件解析后的properties指针类型
//
// 当读取配置文件失败的时候意味着这是一个完全无效且可能会造成后续程序异常的操作，抛出error。
// 读取的配置文件中如果存在“\r\n”或“\n”，抛出error。
// 读取的配置文件中如果发生没有key的情况，抛出error。
func New(propertiesFilePath string) (*Properties, error) {
	fileData, err := kfile.ReadOnce(propertiesFilePath)
	if err != nil {
		return nil, err
	}

	this := &Properties{
		details: make(map[string]string),
	}

	err = kstreams.EachLineOff(string(fileData), func(index int, line string) error {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			return nil
		}
		if kobj.IsEmpty(line) {
			return nil
		}
		key, value := kstr.KV(line)
		if kobj.IsEmpty(key) {
			return errors.New(fmt.Sprintf("not found key by %v (line:%v)", propertiesFilePath, line))

		}
		if strings.Contains(key, "\n") || strings.Contains(key, "\r\n") || strings.Contains(value, "\n") || strings.Contains(value, "\r\n") {
			return errors.New(fmt.Sprintf("Nonconforming \"key\" or \"value\" %v (line:%v)", propertiesFilePath, line))
		}
		this.details[key] = value
		return nil
	})

	return this, err
}
