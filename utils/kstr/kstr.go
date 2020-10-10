// expect：be sure to finish!
// author：KercyLAN
// create at：2019-12-10 11:45

/**
包kstr包含了对字符串处理的相关封装和描述
 */
package kstr

import (
	"strconv"
	"strings"
)

// 检查字符串是否为空字符串，返回bool
func IsEmpty(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}

// 返回str中去除最后一个字符后的字符串。
func RemoveLast(str string) string {
	if len(str) > 0 {
		return str[0:len(str) - 1]
	}
	return str
}

// 返回str中去除第一个字符后的字符串。
func RemoveFirst(str string) string {
	if len(str) > 0 {
		return str[1:]
	}
	return str
}

// 返回str中由多个连续字符char转换为一个字符后的字符串。
func Distinct(str string, char string) string {
	double := char + char
	str = strings.ReplaceAll(str, double, char)
	if strings.Contains(str, double) {
		return Distinct(str, char)
	}
	return str
}

// 返回str是否为大写字母开头的字符串。
func IsUpperPrefix(str string) bool {
	return ((str)[0] >= 65) && ((str)[0] <= 90)
}

// 返回str是否为小写字母开头的字符串。
func IsLowerPrefix(str string) bool {
	return ((str)[0] >= 97) && ((str)[0] <= 122)
}

// 返回将str进行千位分隔符处理后的字符串。
func ThousandsSeparator(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i + 1) * 3] + "," + arr[0][length1-(i + 1) * 3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}

// 返回将str首字符转换为大写后的字符串。
func UpperFirst(str string) string {
	switch len(str) {
	case 0:
		return str
	case 1:
		return strings.ToUpper(string(str[0]))
	default:
		return strings.ToUpper(string(str[0])) + str[1:]
	}
}

// 返回将str首字符转换为小写后的字符串。
func LowerFirst(str string) string {
	switch len(str) {
	case 0:
		return str
	case 1:
		return strings.ToLower(string(str[0]))
	default:
		return strings.ToLower(string(str[0])) + str[1:]
	}
}

// 返回str经过转换后形成的key、value
//
// 这里tag表示使用什么字符串来区分key和value的分隔符。
//
// 默认情况即不传入tag的情况下分隔符为“=”。
func KV(str string, tag ...string) (string, string) {
	tagChar := "="
	if len(tag) > 0 {
		tagChar = tag[0]
	}
	kv := strings.SplitN(str, tagChar, 2)
	if len(kv) < 2 {
		return "", ""
	}
	return kv[0], kv[1]
}

// 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//
// 当字符串为“123,456,789”的时候，返回结果为“123456789”。
//
// 当字符串为“123 456 789”的时候，返回结果为“123456789”。
//
// 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。
func FormatSpeedyInt(numberStr string) (int, error) {
	return strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""))
}

// 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//
// 当字符串为“123,456,789”的时候，返回结果为“123456789”。
//
// 当字符串为“123 456 789”的时候，返回结果为“123456789”。
//
// 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。
func FormatSpeedyInt64(numberStr string) (int64, error) {
	return strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 10, 64)
}

// 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//
// 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
//
// 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
//
// 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。
func FormatSpeedyFloat32(numberStr string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 32)
}

// 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//
// 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
//
// 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
//
// 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。
func FormatSpeedyFloat64(numberStr string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 64)
}