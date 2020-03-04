// expect：be sure to finish!
// author：KercyLAN
// create at：2019-12-10 11:45

package kstr

import (
	"regexp"
	"strings"
)

// HideSensitivity 返回防敏感化后的字符串
//
// 隐藏身份证、邮箱、手机号等敏感信息用*号替代
func HideSensitivity(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			resRs := []rune(str)
			res2 := string(resRs[0:3])
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			rs := []rune(str)
			result = string(rs[0:5]) + "****" + string(rs[7:11])

		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}