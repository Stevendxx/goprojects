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

//
