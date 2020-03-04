// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 23:59

package kcrypto

import (
	"crypto/sha1"
	"encoding/hex"
)

// 对字符串进行SHA1加密并返回其结果。
func SHA1String(str string) string{
	return SHA1Data([]byte(str))
}

// 对字节数组进行SHA1加密并返回其结果。
func SHA1Data(data []byte) string{
	c:=sha1.New()
	c.Write(data)
	return hex.EncodeToString(c.Sum(nil))
}
