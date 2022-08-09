package utils

// 验证客户端登录输入是否合法
func InputIsValid(input int, min int, max int) bool {
	return input >= min && input <= max
}
