// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-5 10:49

/**
包kencoder提供给类封装好的编码器
 */
package kencoder

import "github.com/axgle/mahonia"

// 编码器描述
type Encoder func(data []byte) string

// 解码器声明
var (
	GbkDecoder = mahonia.NewDecoder("gbk")
)

var (
	GbkTo = struct {
		UTF8 Encoder
	}{
		UTF8: func(data []byte) string {
			return GbkDecoder.ConvertString(string(data))
		},
	}
)

// 编码转换
func EncoderTransition(src string,oldEncoder string,newEncoder string) string{
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr:= srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder .Translate([]byte(resStr), true)
	return string(resBytes)
}