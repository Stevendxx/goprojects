package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
)

func main() {
	var str string
	var input int
	var err error

	for {
		utils.ShowClientLoginMenu()

		if _, errInput := fmt.Scanln(&str); errInput != nil {
			fmt.Println("[错误]:输入格式错误,应该是整型[0, 2].")
			continue
		}

		if str == "clear" {
			// Windows cls
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			continue
		} else {
			input, err = strconv.Atoi(str)
			if err != nil {
				fmt.Println("[错误]:输入格式错误,应该是整型[0, 2].")
				continue
			} else if utils.InputIsValid(input, 0, 2) {
				switch input {
				case 1:
					// 登录
				case 2:
					// 注册
				case 0:
					fmt.Println("已退出ICHAT APP...")
					os.Exit(0)
				}
			} else {
				fmt.Println("[错误]:输入范围错误,应该是整型[0, 2].")
				continue
			}
		}

	}
}
