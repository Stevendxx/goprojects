package utils

import (
	"fmt"
)

// 显示客户端登录界面
func ShowClientLoginMenu() {
	fmt.Println(" =============== ICHAT APP ===============")
	fmt.Println("||               1 -> 登录                ||")
	fmt.Println("||               2 -> 注册                ||")
	fmt.Println("||               0 -> 退出                ||")
	fmt.Println(" =========================================")
	fmt.Print("input> ")
}

// 显示用户界面
func ShowUserMenu(name string) {
	fmt.Printf(" 当前用户: %s 状态: 在线\n", name)
	fmt.Println(" =============== ICHAT APP ===============")
	fmt.Println("||              1 -> 在线用户             ||")
	fmt.Println("||              2 -> 发送消息             ||")
	fmt.Println("||              3 -> 我的好友             ||")
	fmt.Println("||              4 -> 我的消息             ||")
	fmt.Println("||              0 -> 退出软件             ||")
	fmt.Println(" =========================================")
	fmt.Print("input> ")
}
