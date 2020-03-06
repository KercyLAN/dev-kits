// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-6 0:06

package main

import (
	"runtime"
)

func main()  {
	b := make([]byte, 7 * 1024 * 1024 * 1024 + 4200)
	for i:= 0; i< len(b); i++ {
		b[i] = byte('1')
	}
	runtime.GC()
}
