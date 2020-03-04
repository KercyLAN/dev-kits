// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 23:59

package kcrypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// 对字符串进行SHA256加密并返回其结果。
func SHA256String(str string) string{
	return SHA256Data([]byte(str))
}

// 对字节数组进行SHA256加密并返回其结果。
func SHA256Data(data []byte) string{
	bytes := sha256.Sum256(data)
	return hex.EncodeToString(bytes[:])
}
