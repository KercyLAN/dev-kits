// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 23:57

package kcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

// 对字符串进行MD5加密并返回其结果。
func MD5String(str string) string {
	return MD5Data([]byte(str))
}

// 对字节数组进行MD5加密并返回其结果。
func MD5Data(data []byte) string {
	c := md5.New()
	c.Write(data)
	return hex.EncodeToString(c.Sum(nil))
}

