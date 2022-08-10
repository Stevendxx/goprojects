package model

type ClientUser struct {
	id     uint
	name   string
	age    byte
	gender rune

	// 好友列表
	friends []ClientUser
}
