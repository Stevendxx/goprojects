package model

const (
	LoginMsgType  = "LoginMessage"
	LogonMsgType  = "LogonMessage"
	LoginCodeType = "LoginCode"
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
type LogonMessage struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

// 服务器返回码
type LoginCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
