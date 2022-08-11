package model

const (
	LoginMsgType   = "LoginMessage"
	SignUpMsgType  = "SignUpMessage"
	LoginCodeType  = "LoginCode"
	SignUpCodeType = "SignUpCode"
)

// 交互报文
type Message struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	Length int    `json:"length"`
}

// 登录报文
type LoginMessage struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

// 注册报文
type SignUpMessage struct {
	SignUpUser User `json:"user"`
}

// 登录返回码
/*
 - 200 成功
 - 500 失败
 - 123 密码错误
*/
type LoginCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 注册返回码
/*
 - 200 成功
 - 500 失败
 - 400 写存失败
*/
type SignUpCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
