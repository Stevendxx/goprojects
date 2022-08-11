package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
)

func main() {
	var str string
	var input int
	var err error
	//var user = &model.User{}

	for {
		utils.ShowClientLoginMenu()

		if _, errInput := fmt.Scanln(&str); errInput != nil {
			fmt.Printf("[错误]:输入发生错误,应该是整型[0, 2].\n\n")
			continue
		}

		if str == "clear" {
			utils.ClearScreen()
			continue
		} else {
			input, err = strconv.Atoi(str)
			if err != nil {
				fmt.Printf("[错误]:输入格式错误,应该是整型[0, 2].\n\n")
				continue
			} else if utils.InputIsValid(input, 0, 2) {
				switch input {
				case 1:
					// 登录
					LoginMenu()
				case 2:
					// 注册
					SignUpMenu()
				case 0:
					fmt.Println("已退出ICHAT APP...")
					os.Exit(0)
				}
			} else {
				fmt.Printf("[错误]:输入范围错误,应该是整型[0, 2].\n\n")
				continue
			}
		}

	}
}
