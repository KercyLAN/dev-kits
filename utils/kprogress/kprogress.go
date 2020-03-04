// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 10:35

package kprogress

import (
	"fmt"
	"strconv"
)

// 根据now和gross通过百分率计算进度保留两位小数并反馈到hook中。
func Progress(now int64, gross int64, hook func(now int64, gross int64, progress float64, finish bool)) {
	progress, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(now) / float64(gross) * 100), 64)
	finish := false
	if now >= gross{
		finish = true
	}
	hook(now, gross, progress, finish)
}