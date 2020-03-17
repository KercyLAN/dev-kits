// expect：be sure to finish!
// author：KercyLAN
// create at：2019-12-10 13:38

package kstreams

import (
	"errors"
	"reflect"
	"sort"
)

// 对传入的eachMap进行排序后再进行遍历
//
// 确保eachFunc传入的类型为func，并且其拥有两个入参，分别对应eachMap的Key类型和Value类型。
//
// 在EachMapSort对eachMap排序之后会进行一次遍历，并且将Key和Value分别传入eachFunc这个函数里进行调用。
//
// 使用示例如下：
//	mInt := map[int]string {
//		3: "line 3",
//		1: "line 1",
//		2: "line 2",
//	}
//	EachMapSort(mInt, func(key int, value string) {
//		t.Log(key, value)
//	})
//	t.Log("------------------------------------")
//	mString := map[string]string {
//		"3": "line 3",
//		"1": "line 1",
//		"b": "line b",
//		"2": "line 2",
//		"a": "line a",
//	}
//	EachMapSort(mString, func(key string, value string) {
//		t.Log(key, value)
//	})
func EachMapSort(eachMap interface{}, eachFunc interface{})  {
	eachMapValue := reflect.ValueOf(eachMap)
	eachFuncValue := reflect.ValueOf(eachFunc)
	eachMapType := eachMapValue.Type()
	eachFuncType := eachFuncValue.Type()
	if eachMapValue.Kind() != reflect.Map {
		panic(errors.New("kstreams.EachMap failed. parameter \"eachMap\" type must is map[...]...{}"))
	}
	if eachFuncValue.Kind() != reflect.Func {
		panic(errors.New("kstreams.EachMap failed. parameter \"eachFunc\" type must is func(key ..., value ...)"))
	}
	if eachFuncType.NumIn() != 2 {
		panic(errors.New("kstreams.EachMap failed. \"eachFunc\" input parameter count must is 2"))
	}
	if eachFuncType.In(0).Kind() != eachMapType.Key().Kind() {
		panic(errors.New("kstreams.EachMap failed. \"eachFunc\" input parameter 1 type not equal of \"eachMap\" key"))
	}
	if eachFuncType.In(1).Kind() != eachMapType.Elem().Kind() {
		panic(errors.New("kstreams.EachMap failed. \"eachFunc\" input parameter 2 type not equal of \"eachMap\" value"))
	}

	// 对key进行排序
	// 获取排序后map的key和value，作为参数调用eachFunc即可
	switch eachMapType.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keys := make([]int, 0)
		keysMap := map[int]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, int(value.Int()))
			keysMap[int(value.Int())] = value
		}
		sort.Ints(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.Float64, reflect.Float32:
		keys := make([]float64, 0)
		keysMap := map[float64]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, float64(value.Float()))
			keysMap[float64(value.Float())] = value
		}
		sort.Float64s(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.String:
		keys := make([]string, 0)
		keysMap := map[string]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, value.String())
			keysMap[value.String()] = value
		}
		sort.Strings(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	default:
		panic(errors.New("\"eachMap\" key type must is int or float or string"))
	}
}
