package hanzi

import (
	"github.com/KercyLAN/dev-kits/components/hanzi/pinyin"
	"github.com/KercyLAN/dev-kits/utils/kstr"
	"strings"
)


// 传入一串文本，转换为汉字信息
func ToHanzi(word string) *hanzi {
	hz := &hanzi{
		char: word,
		rhyme: make([]string, 0),
	}
	// 取到每一个字符转换后的Pinyin
	for _, py := range pinyin.Gains(word) {
		if len(py.Finals()) > 0 {
			hz.rhyme = append(hz.rhyme, py.Finals()[0])
		}
		hz.pinyin = append(hz.pinyin, py)
	}
	return hz
}

// 传入一段字符串,获取字符串中包含的所有词语
// todo 如果传入文字过长，效率较低，待优化
func TakeWordsFull(str string) []*hanzi {
	result := make([]*hanzi, 0)

	zi := strings.Split(str, "")
	tempZi := ""
	for _, val := range zi {
		tempZi += val
		if !kstr.IsEmpty(chineseDir[tempZi]) {
			hz := &hanzi{
				char: tempZi,
				rhyme: make([]string, 0),
			}
			for _, py := range pinyin.Gains(tempZi) {
				if len(py.Finals()) > 0 {
					hz.rhyme = append(hz.rhyme, py.Finals()[0])
				}
				hz.pinyin = append(hz.pinyin, py)
			}
			result = append(result, hz)
			tempZi = ""
		}
	}
	if len(str) > 0 {
		result = append(result, TakeWordsFull(string([]byte(str)[1:len([]byte(str))]))...)
	}
	// 去重
	endStr := ""
	end := make([]*hanzi, 0)
	for _, r := range result {
		if !strings.Contains(endStr, r.char) {
			end = append(end, r)
			endStr += r.char
		}
	}
	return end
}

// 押韵
func Rhyme(word string, dir []*hanzi) []*hanzi {
	hz := ToHanzi(word)
	result := make([]*hanzi, 0)
	addStr := ""
	for _, val := range dir {
		r1 := ""
		r2 := ""
		for _, value := range hz.rhyme {
			r1 += value
		}
		for _, value := range val.rhyme {
			r2 += value
			if r1 == r2 {
				if !strings.Contains(addStr, val.String()) {
					addStr += val.String()
					result = append(result, val)
					break
				}
			}
		}
	}
	return result
}