// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-28 19:34

package ktime

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Interval 返回了指定时间戳与当前时间的间隔
//
// 使用传入的时间和当前的时间进行计算换算，将结果体现为几年前、几天前、几小时前、几分钟前、几秒前
func Interval(timestamp int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - timestamp
	if ct <= 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = fmt.Sprint(tempStr, unit[i])
		}
		break
	}
	return res
}

