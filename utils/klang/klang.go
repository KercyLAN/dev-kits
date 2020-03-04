// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 22:56

/**
包klang包含了对国际化支持的封装及实现
 */
package klang

import (
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kproperties"
)

// 国际化语言代码描述
type Local string

// Lang 国际化语言实现
type Lang struct {
	propertiesDefault			kproperties.PropertiesInterface				// 默认语言文件信息的properties实现
	properties					map[Local]kproperties.PropertiesInterface	// 语言文件信息的properties实现集合
	bundle						string										// 对应语言文件名
	bundleDirPath 				string										// bundle所在的文件夹路径
	local						Local										// 地区
	encoder						func(str string) string						// 编码器

	sweepersThresholdValue		int											// 清理器阈值，当满足这个值的时候才会触发清理器进行清理
	sweepers					func(lang *Lang)							// 清理器，当该国际化中支持快速切换的语言数量达到阈值将会触发该函数
}

// 设置清理器
//
// 清理器用来解决在语言频繁切换的情况下导致过多的语言沉积造成内容开销过大的问题。
//
// 当缓存的语言数量达到sweepersThresholdValue的值的时候，则会调用清理函数进行清理。
func (slf *Lang) SetSweepers(sweepersHandler func(lang *Lang)) {
	slf.sweepers = sweepersHandler
	if slf.sweepersThresholdValue == 0 {
		slf.sweepersThresholdValue = 5
	}
}

// 设置清理器触发阈值
//
// 当传入的value小于0的时候，则会将value设置为0，表示不触发清理器。
func (slf *Lang) SetSweepersThresholdValue(value int) {
	if value < 0 {
		value = 0
	}
	slf.sweepersThresholdValue = value
}

// 将国际化中支持快速切换的语言全部清空仅保留当前使用的语言。
func (slf *Lang) Reset() {
	nowProperties := slf.properties[slf.local]
	slf.properties = map[Local]kproperties.PropertiesInterface{
		slf.local: nowProperties,
	}
}

// 返回当前国际化中支持快速切换的语言数量
//
// 快速切换即指之前该Lang已经使用过该语言，仅需要内存中直接获取即可，而非从文件读取。
func (slf *Lang) FastSwitchSize() int {
	return len(slf.properties)
}

// 返回当前国际化中支持快速切换的语言
//
// 快速切换即指之前该Lang已经使用过该语言，仅需要内存中直接获取即可，而非从文件读取。
func (slf *Lang) FastSwitchSizeLocals() []Local {
	locals := make([]Local, 0)
	for local := range slf.properties{
		locals = append(locals, local)
	}
	return locals
}

// 返回当前特定地区语言文件中key对应的value
//
// 指定语言文件中缺少这部分翻译，则从默认语言文件中提取，
// 若是均找不到key则返回长度为0的字符串。
func (slf *Lang) Get(key string) string {
	var value string
	nowProperties := slf.properties[slf.local]
	if nowProperties.HasKey(key) {
		value = nowProperties.GetString(key)
	}else {
		value = slf.propertiesDefault.GetString(key)
	}
	if slf.encoder != nil {
		return slf.encoder(value)
	}
	return value
}

// 设置编码器
//
// 编码器用于将读取到的语言内容转换为特定编码。
//
// 如果存在编码器，那么通过Get获取到的值将会通过编码器进行转码后再返回。
// 如果不设置编码器，那么默认编码器为klang.EncoderGbkUtf8。
func (slf *Lang) SetEncoder(encoderHandler func(str string) string) {
	slf.encoder = encoderHandler
}

// 设置当前Lang的地区信息
//
// 在设置前会查找之前是否已经有过该语言的设置，如果有直接调整即可，否则反复读取文件可能造成性能消耗。
//
// 当找不到特定地区语言文件或文件内容存在异常则不会进行设置，而是返回一个error。
func (slf *Lang) SetLocal(local Local) error {
	if _, ok := slf.properties[local]; ok {
		slf.local = local
		return nil
	} else {
		if tryProperties, err := kproperties.New(productionPath(slf.bundle, slf.bundleDirPath, local)); err != nil {
			return err
		}else {
			slf.local = local
			slf.properties[local] = tryProperties
			if len(slf.properties) >= slf.sweepersThresholdValue {
				slf.sweepers(slf)
			}
			return nil
		}
	}
}

// 获取当前Lang的地区信息。
func (slf *Lang) GetLocal() Local {
	return slf.local
}

// 构建一个指定地区的国际化语言实例
//
// local来源于klang.xxx。
//
// bundle表示了对应语言文件名，应该与“bundleName.local.properties”类似，如“message_zh-CN.properties”。
//
// bundleDirPath表示了bundle所在的文件夹路径。
//
// 当找不到特定地区语言文件或文件内容存在异常则会发生panic。
func New(bundle string, bundleDirPath string, local Local) *Lang {
	tryProperties, err := kproperties.New(productionPath(bundle, bundleDirPath, local))
	if err != nil {
		panic(err)
	}

	defaultProperties, err := kproperties.New(productionPath(bundle, bundleDirPath, LOCAL_NONE))
	if err != nil {
		panic(err)
	}

	this := &Lang{
		propertiesDefault:defaultProperties,
		properties: map[Local]kproperties.PropertiesInterface{
			local: tryProperties,
		},
		bundle:        bundle,
		bundleDirPath: bundleDirPath,
		local:         local,
		encoder:       EncoderGbkUtf8,

		sweepersThresholdValue:0,
		sweepers: func(lang *Lang) {
			// 表示默认不做任何清理
		},
	}


	return this
}

// 内部生产一个符合Local的properties地址并返回
//
// 如果local为NONE的话，那么就应该取默认的语言properties，其应该缺少local片段，如“message.properties”。
func productionPath(bundle, bundleDirPath string, local Local) string {
	if local != LOCAL_NONE {
		return fmt.Sprintf("%s\\%s_%s.properties", bundleDirPath, bundle, local)
	}else {
		return fmt.Sprintf("%s\\%s.properties", bundleDirPath, bundle)
	}
}