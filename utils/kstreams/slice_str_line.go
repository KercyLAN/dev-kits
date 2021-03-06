// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 21:28

/**
包kstreams包含了对数据处理的相关封装及描述
 */
package kstreams

import "strings"

// 对传入的eachString进行按行切片后再进行遍历
//
// 该函数会预先对“\r\n”进行处理替换为“\n”。
//
// 在遍历到每一行的时候会将结果index和line作为入参传入eachFunc中进行调用。
//
// index表示了当前行的行号（由0开始），line表示了当前行的内容。
func EachLine(eachString string, eachFunc func(index int, line string)) {
	formatStr := strings.ReplaceAll(eachString, "\r\n", "\n")
	for index, line := range strings.Split(formatStr, "\n") {
		eachFunc(index, line)
	}
}

// 对传入的eachString进行按行切片后再进行可中断的遍历
//
// 该函数会预先对“\r\n”进行处理替换为“\n”。
//
// 在遍历到每一行的时候会将结果index和line作为入参传入eachFunc中进行调用。
//
// index表示了当前行的行号（由0开始），line表示了当前行的内容。
//
// 当eachFunc的返回值不为nil时，则中断遍历，并返回error。
func EachLineOff(eachString string, eachFunc func(index int, line string) error) error {
	formatStr := strings.ReplaceAll(eachString, "\r\n", "\n")
	for index, line := range strings.Split(formatStr, "\n") {
		err := eachFunc(index, line)
		if err != nil {
			return err
		}
	}
	return nil
}