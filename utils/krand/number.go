package krand

import "math/rand"

// 返回一个介于min和max之间的int64类型的随机数。
func Int64(min int64, max int64) int64 {
	return min + rand.Int63n(max - min)
}

// 返回一个介于min和max之间的的int类型的随机数。
func Int(min int, max int) int {
	return int(Int64(int64(min), int64(max)))
}

