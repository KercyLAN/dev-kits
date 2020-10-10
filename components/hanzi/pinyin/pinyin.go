package pinyin

import (
	"strings"
	"unicode/utf8"
)

// Pinyin 拼音结构
type Pinyin struct {
	source 			string 			// 源文本
	result 			string			// 拼音
	initial			string 			// 声母
	finals 			[]string 		// 韵母
	tone			int 			// 声调
	err 			error			// 解析错误
}

// String 获取不带声调的完整拼音
func (slf *Pinyin) String() string {
	if slf.result == " " {
		return " "
	}
	if slf.err != nil {
		return slf.result
	}
	output := make([]rune, utf8.RuneCountInString(slf.result))
	count := 0
	for _, t := range slf.result {
		neutral, found := rhymeTonesMapping[t]
		if found {
			output[count] = neutral
		} else {
			output[count] = t
		}
		count++
	}
	return string(output)
}

// StringTone 获取完整的拼音
func (slf *Pinyin) StringTone() string {
	if slf.result == " " {
		return " "
	}
	return slf.result
}

// getFinals 获取韵母
func (slf *Pinyin) getFinals() {
	for _, value := range templateRhyme {
		if strings.Contains(slf.String(), value) {
			// 判断已有的韵母中是否包含了这个内容
			isAdd := true
			for _, final := range slf.finals {
				if strings.Contains(final, value) {
					isAdd = false
				}
			}
			if isAdd {
				slf.finals = append(slf.finals, value)
			}
		}
	}
}

// getInitials 获取声母
func (slf *Pinyin) getInitials() {
	temp := strings.Split(slf.result, "")
	if len(temp) < 2 {
		return
	}
	// 检测zh ch sh
	for _, initial := range templateInitials {
		if initial == temp[0] + temp[1] {
			slf.initial = initial
			return
		}
	}
	// 其他
	for _, initial := range templateInitials {
		if initial == temp[0] {
			slf.initial = initial
			return
		}
	}
}


func (slf *Pinyin) Tone() int {
	return slf.tone
}

func (slf *Pinyin) Source() string {
	return slf.source
}

func (slf *Pinyin) Initial() string {
	return slf.initial
}

func (slf *Pinyin) Finals() []string {
	return slf.finals
}

func (slf *Pinyin) Err() error {
	return slf.err
}
