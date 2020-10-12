// Package brush 实现了同时兼容Linux及Windows的控制台彩色输出封装
package brush

import (
	"fmt"
	"os"
	"strings"
	"github.com/shiena/ansicolor"
)
// brush 笔刷结构定义
type brush struct {
	background		string	// 背景色
	foreground 		string	// 前景色
	Style			string	// 样式
}

// Print 兼容Linux和Windows控制台的打印
func Print(text string)  {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	w.Write([]byte(text))
}

// Printf 兼容Linux和Windows控制台的打印
func Printf(text string)  {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	w.Write([]byte(text + "\n"))
}

// BuildBrush 构建一个笔刷
func BuildBrush() *brush {
	return &brush{
		background: "40",
		foreground: "37",
		Style:		"",
	}
}

// Clear 清除笔刷附加给文字的色彩信息
func Clear(text string) string {
	var result string
	// 根据头部分隔为单独的片段
	// x;x;xm文本\033[0m, x;x;xm文本\033[0m, x;x;xm文本\033[0m
	textTemp := strings.ReplaceAll(text, "\033[", "\\033[")
	textTemp = strings.ReplaceAll(textTemp, "\033[0m", "\\033[0m")
	slice := strings.Split(textTemp, "\\033[")
	for _, value := range slice {
		if strings.Contains(value, "m") {
			colorAndText := strings.Split(value, "\\033[0m")[0]
			colorAndTextSlice := strings.SplitN(colorAndText, "m", 2)
			if len(colorAndTextSlice) == 2 {
				result += colorAndTextSlice[1]
			}
			continue
		}
		result += value
	}
	return result
}


// Go 上色
func (slf *brush) Go(text string) string {
	return fmt.Sprint("\033[", slf.background, ";", slf.foreground, ";", slf.Style, "m", text, "\033[0m")
}


// StyleNone 样式无
func (slf *brush) StyleNone() *brush {
	slf.Style = "0"
	return slf
}

// StyleHighlight 样式好亮
func (slf *brush) StyleHighlight() *brush {
	slf.Style = "1"
	return slf
}

// StyleUnderline 样式下划线
func (slf *brush) StyleUnderline() *brush {
	slf.Style = "4"
	return slf
}

// StyleFlicker 样式闪烁
func (slf *brush) StyleFlicker() *brush {
	slf.Style = "5"
	return slf
}

// StyleInverse 样式反白
func (slf *brush) StyleInverse() *brush {
	slf.Style = "7"
	return slf
}

// StyleInvisible 样式不可见
func (slf *brush) StyleInvisible() *brush {
	slf.Style = "8"
	return slf
}

// FgBlack 前景黑色
func (slf *brush) FgBlack() *brush {
	slf.foreground= "30"
	return slf
}

// FgRed 前景红色
func (slf *brush) FgRed() *brush {
	slf.foreground= "31"
	return slf
}

// FgGreen 前景绿色
func (slf *brush) FgGreen() *brush {
	slf.foreground= "3"
	return slf
}

// FgYellow 前景黄色
func (slf *brush) FgYellow() *brush {
	slf.foreground= "33"
	return slf
}

// FgBlue 前景蓝色
func (slf *brush) FgBlue() *brush {
	slf.foreground= "34"
	return slf
}

// FgPurple 前景紫色
func (slf *brush) FgPurple() *brush {
	slf.foreground= "35"
	return slf
}

// FgCyan 前景青色
func (slf *brush) FgCyan() *brush {
	slf.foreground= "36"
	return slf
}

// FgWhite 前景白色
func (slf *brush) FgWhite() *brush {
	slf.foreground= "37"
	return slf
}

// BgBlack 背景黑色
func (slf *brush) BgBlack() *brush {
	slf.background= "40"
	return slf
}

// BgRed 背景红色
func (slf *brush) BgRed() *brush {
	slf.background= "41"
	return slf
}

// BgGreen 背景绿色
func (slf *brush) BgGreen() *brush {
	slf.background= "42"
	return slf
}

// BgYellow 背景黄色
func (slf *brush) BgYellow() *brush {
	slf.background= "43"
	return slf
}

// BgBlue 背景蓝色
func (slf *brush) BgBlue() *brush {
	slf.background= "44"
	return slf
}

// BgPurple 背景紫色
func (slf *brush) BgPurple() *brush {
	slf.background= "45"
	return slf
}

// BgCyan 背景青色
func (slf *brush) BgCyan() *brush {
	slf.background= "46"
	return slf
}

// BgWhite 背景白色
func (slf *brush) BgWhite() *brush {
	slf.background= "47"
	return slf
}