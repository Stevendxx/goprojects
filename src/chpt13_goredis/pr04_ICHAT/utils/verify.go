package utils

// 验证输入是否合法
func InputIsValid(input int, min int, max int) bool {
	return input >= min && input <= max
}
