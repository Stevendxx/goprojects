package model

type User struct {
	Id     int    `json:"userId"`
	Name   string `json:"userName"`
	Pwd    string `json:"userPwd"`
	Age    byte   `json:"age"`
	Gender rune   `json:"gender"`
	Email  string `json:"email"`
	State  string `json:"state"`

	// 好友列表
	Friends []User `json:"friends"`
}
