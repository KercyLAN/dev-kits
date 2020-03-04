// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 20:04

/**
包kobj包含了所有对对象（object）的描述及封装
 */
package kobj

import "reflect"

// 返回i是否为空。
func IsEmpty(i interface{}) bool {
	switch reflect.ValueOf(i).Kind() {
	case reflect.String:
		if len(i.(string)) == 0 {
			return true
		}
	case reflect.Int:
		return i.(int) == 0
	case reflect.Int8:
		return i.(int8) == 0
	case reflect.Int16:
		return i.(int16) == 0
	case reflect.Int32:
		return i.(int32) == 0
	case reflect.Int64:
		return i.(int64) == 0
	case reflect.Uint:
		return i.(uint) == 0
	case reflect.Uint8:
		return i.(uint8) == 0
	case reflect.Uint16:
		return i.(uint16) == 0
	case reflect.Uint32:
		return i.(uint32) == 0
	case reflect.Uint64:
		return i.(uint64) == 0
	case reflect.Float32:
		return i.(float32) == 0
	case reflect.Float64:
		return i.(float64) == 0
	case reflect.Bool:
		return i.(bool) == false
	case reflect.Struct, reflect.Ptr, reflect.Slice, reflect.Map:
		return reflect.ValueOf(i).IsNil()
	}
	return i == nil
}

// 返回传入的所有内容是否为空
//
// 等同于循环调用IsEmpty进行检查，当一个不为空则全盘推翻。
func IsAllEmpty(i ...interface{}) bool {
	for _, value := range i {
		if !IsEmpty(value) {
			return false
		}
	}
	return true
}
