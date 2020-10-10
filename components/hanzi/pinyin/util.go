package pinyin

import (
	"errors"
	"github.com/KercyLAN/dev-kits/utils/kstr"
	"regexp"
	"strings"
)

type converMode int
var(
	modeString converMode = 1			// 不带音调
	modeStringTone converMode = 2		// 携带音调
)

// 将一组Pinyin转换为不带声调的字符串
func ToString(pinyins []*Pinyin, split ...string) string {
	return conver(modeString, pinyins, split...)
}

// 将一组Pinyin转换为带声调的字符串
func ToStringTone(pinyins []*Pinyin, split ...string) string {
	return conver(modeStringTone, pinyins, split...)
}

// 获取单个文字拼音
func Gain(char string) *Pinyin {
	py := &Pinyin{source:char}
	// 验证传入值
	if !check(char, py) {
		return py
	}
	// 取到对应值
	pyStr := pinyinDictionary[[]rune(char)[0]]
	if kstr.IsEmpty(pyStr) {
		py.result = char
		py.err = errors.New("Incorrect characters were typed")
		return py
	}
	// 获取声调
	for _, t := range pyStr {
		_, found := rhymeTonesMapping[t]
		// 如果找到则赋值，否则取最后一个字母
		if found {
			py.tone = rhymeTonesOrderMapping[t]
		}
	}
	// 获取声母韵母
	py.result = pyStr
	py.getFinals()
	py.getInitials()

	return py
}

// 获取多个文字拼音
func Gains(str string) []*Pinyin {
	result := make([]*Pinyin, 0)
	slice := strings.Split(str, "")
	nowP := 0
	for position, char := range slice {
		word := char
		if nowP > position {
			continue
		}
		nowP = position
		// 检测是否为英文，如果是。则和下一个拼接
		match, err := regexp.MatchString("^[A-Za-z']*$",char)
		if err != nil {
			return result
		}
		if match {
			// 拼接连续的英文或者 ' 符号
			for i := position + 1; i < len(slice); i++ {
				m, e := regexp.MatchString("^[A-Za-z']*$", slice[i])
				if e != nil {
					return result
				}
				if m {
					word += slice[i]
				}else {
					break
				}
				nowP = i + 1
			}
		}
		if kstr.IsEmpty(strings.TrimSpace(word)) {
			result = append(result, Gain(" "))
			continue
		}
		result = append(result, Gain(word))
	}
	return result
}

// 快捷转换
func conver(mode converMode, pinyins []*Pinyin, split ...string) string {
	sp := " "
	result := ""
	if len(split) != 0 {
		sp = split[0]
	}
	for _, pinyin := range pinyins {
		if pinyin.result == sp {
			switch mode {
			case modeString:
				result += pinyin.String()
			case modeStringTone:
				result += pinyin.StringTone()
			}
		}else {
			switch mode {
			case modeString:
				result += pinyin.String() + sp
			case modeStringTone:
				result += pinyin.StringTone() + sp
			}
		}
	}
	return kstr.RemoveLast(result)
}

// 验证待转拼音的字符
func check(char string, pinyin *Pinyin) bool {
	// 检测为空格
	if kstr.Distinct(char, " ") == " " {
		pinyin.result = " "
		return false
	}
	// 检测是否为英文单词
	match, err := regexp.MatchString("^[A-Za-z']*$",char)
	if err != nil {
		pinyin.result = char
		pinyin.err = errors.New("Incorrect characters were typed")
		return false
	}
	if match {
		pinyin.result = char
		return false
	}
	// 检测长度
	if len(char) != 3 {
		pinyin.result = char
		pinyin.err = errors.New("Only one character is allowed")
		return false
	}
	return true
}