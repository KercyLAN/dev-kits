// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-27 21:19

/**
包kruntime包含了对程序运行时的相关封装及描述
 */
package kruntime

import (
	"os"
	"path/filepath"
	"runtime"
)


var (
	pathRunDir string													// 程序运行的绝对路径
	pathRun string														// 应用程序文件的绝对路径
	pathWork string 													// 项目工作目录
)

func init() {
	var err error
	pathRunDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	pathRun = os.Args[0]
	pathWork, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

// 返回当前进程可用的逻辑cpu数量
//
// 通过查询操作系统来检查可用的cpu集，调用了go lib中runtime.NumCPU函数。
func NumCPU() int {
	return runtime.NumCPU()
}

// 返回当前程序运行目录的绝对路径
//
// 默认kercylan-lib库的kruntime包在被调用之处便通过初始化获取了当前应用程序运行的根目录绝对路径。
//
// 由于os.Args[0]获取的是应用程序的运行路径，那么通过filepath.Dir获取到应用程序的运行目录后通过filepath.Abs将其转换确保其为一个绝对路径。
func PathRunDir() string {
	return pathRunDir
}

// 返回当前运行的应用程序的绝对路径
//
// os.Args[0]得到的值直接表示了应用程序的运行路径，并且这是一个绝对路径，等同于os.Args[0]。
func PathRunExe() string {
	return pathRun
}

// 返回当前运行的应用程序的工作路径
//
// 等同于os.Getwd。
func PathWork() string {
	return pathWork
}
