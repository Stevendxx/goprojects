package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
)

// 登录验证
func LoginVerfiy(name, pwd string) (conn net.Conn, msg model.Message, err error) {
	if name == "" && pwd == "" {
		err = model.ErrorUserNamePwdNil
		return
	} else if name == "" {
		err = model.ErrorUserNameNil
		return
	} else if pwd == "" {
		err = model.ErrorUserPwdNil
		return
	}

	// 1. 请求连接
	conn, err = utils.GetServerConn()
	if err != nil {
		return
	}

	// 2. 登录信息
	var loginMsg = model.LoginMessage{
		UserId:   0,
		UserName: name,
		UserPwd:  pwd,
	}

	msg, err = utils.CreateMsg(model.LoginMsgType, loginMsg)
	if err != nil {
		return
	}

	// 5. 发送报文至服务器
	errWrite := utils.WrtieMsg(conn, msg)
	if errWrite != nil {
		if errWrite == io.EOF {
			fmt.Println("[客户端]:服务器关闭...")
		} else {
			fmt.Println("[错误]:发送报文出现错误.")
			fmt.Println(errWrite)
		}
		return
	}

	// 读取服务器返回报文
	returnMsg, errRead := utils.ReadMsg(conn)
	if errRead != nil {
		if errRead == io.EOF {
			fmt.Println("[客户端]:服务器关闭...")
		} else {
			fmt.Println("[错误]:读取报文出现错误.")
			fmt.Println(errRead)
		}
		return
	}
	msg = returnMsg
	return
}

// 注册验证
func SignUpVerfiy(user model.User) (msg model.Message, err error) {
	// 1. 请求连接
	conn, err := utils.GetServerConn()
	if err != nil {
		return
	}
	defer conn.Close()

	// 2. 登录信息
	var signUpMsg = model.SignUpMessage{
		SignUpUser: user,
	}

	msg, err = utils.CreateMsg(model.SignUpMsgType, signUpMsg)
	if err != nil {
		return
	}

	// 5. 发送报文至服务器
	errWrite := utils.WrtieMsg(conn, msg)
	if errWrite != nil {
		if errWrite == io.EOF {
			fmt.Println("[客户端]:服务器关闭...")
		} else {
			fmt.Println("[错误]:发送报文出现错误.")
			fmt.Println(errWrite)
		}
		return
	}

	// 读取服务器返回报文
	returnMsg, errRead := utils.ReadMsg(conn)
	if errRead != nil {
		if errRead == io.EOF {
			fmt.Println("[客户端]:服务器关闭...")
		} else {
			fmt.Println("[错误]:读取报文出现错误.")
			fmt.Println(errRead)
		}
		return
	}
	msg = returnMsg
	return
}

// 处理登录返回报文
func HandleLoginReturnMsg(msg model.Message) (err error) {

	var loginCode model.LoginCode
	err = json.Unmarshal([]byte(msg.Data), &loginCode)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}
	//fmt.Println(loginCode.Code)
	if loginCode.Code != 200 {
		err = errors.New(loginCode.Message)
		return
	}
	return
}

// 处理注册返回报文
func HandleSignUpReturnMsg(msg model.Message) (err error) {

	var signUpCode model.SignUpCode
	err = json.Unmarshal([]byte(msg.Data), &signUpCode)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}

	if signUpCode.Code != 200 {
		err = errors.New(signUpCode.Message)
		return
	}
	return
}

// 用户主程序逻辑处理
func UserMain(conn net.Conn, name, pwd string) {
	defer conn.Close()
	var str string
	var input int
	var err error

	// 上线时接受来自群聊或好友的消息
	go AcceptMsg(conn)

	for {
		utils.ShowUserMenu(name)

		if _, errInput := fmt.Scanln(&str); errInput != nil {
			fmt.Printf("[错误]:输入发生错误,应该是整型[0, 4].\n\n")
			continue
		}

		if str == "clear" {
			utils.ClearScreen()
			continue
		} else {
			input, err = strconv.Atoi(str)
			if err != nil {
				fmt.Printf("[错误]:输入格式错误,应该是整型[0, 4].\n\n")
				continue
			} else if utils.InputIsValid(input, 0, 4) {
				switch input {
				case 1:
					// 在线用户
				case 2:
					// 发送消息
				case 3:
					// 我的好友
				case 4:
					// 我的消息
				case 0:
					// 退出软件
					fmt.Print("正在退出登录")
					utils.ProgressBar(3)
					return
				}
			} else {
				fmt.Printf("[错误]:输入范围错误,应该是整型[0, 4].\n\n")
				continue
			}
		}

	}
}

// 接受其他用户的消息
func AcceptMsg(conn net.Conn) {
	for {
		msg, errRead := utils.ReadMsg(conn)
		if errRead != nil {
			if errRead == io.EOF {
				fmt.Println("[客户端]:服务器关闭...")
			} /*else {
				fmt.Println("[错误]:读取报文出现错误.")
				fmt.Println(errRead)
			}*/
			return
		}
		fmt.Println(msg)
	}
}
