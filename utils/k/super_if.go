package k

// 实现了if的三目表达式。
func If(expression bool, satisfy, dissatisfaction interface{}) interface{} {
	if expression {
		return satisfy
	}
	return dissatisfaction
}