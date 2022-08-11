package utils

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
)

// 报文传输

type Transfer struct {
}

// 获取服务器连接
func GetServerConn() (conn net.Conn, err error) {
	network := "tcp"
	address := "localhost:8991"
	conn, errDial := net.Dial(network, address)
	if errDial != nil {
		fmt.Println("[错误]:请求连接错误,请检查网络方式和主机IP及其端口.")
		fmt.Println(errDial)
		err = errDial
		return
	}
	return
}

// 生成报文
func CreateMsg(msgType string, dataMsg interface{}) (msg model.Message, err error) {
	// 1. 元信息序列化
	data, errJson := json.Marshal(dataMsg)
	if errJson != nil {
		fmt.Println("[错误]: 登录信息序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	// 2. 通过序列化后的登录信息构建报文
	msg = model.Message{
		Type:   msgType,
		Data:   string(data),
		Length: len(data),
	}
	return
}

// 解析报文

// 读取报文
func ReadMsg(conn net.Conn) (msg model.Message, err error) {
	var b = make([]byte, 10240)
	cnt, errRead := conn.Read(b)
	if errRead != nil {
		fmt.Println("[错误]:报文读取错误.")
		fmt.Println(errRead)
		err = errRead
		return
	}

	errUnjson := json.Unmarshal(b[:cnt], &msg)
	if errUnjson != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(errUnjson)
		err = errUnjson
		return
	}
	return
}

// 写入报文
func WrtieMsg(conn net.Conn, msg model.Message) (err error) {
	data, errJson := json.Marshal(msg)
	if errJson != nil {
		fmt.Println("[错误]: 报文序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	_, errWrtie := conn.Write([]byte(data))
	if errWrtie != nil {
		fmt.Println("[错误]:发送报文出现错误.")
		fmt.Println(errWrtie)
		err = errWrtie
		return
	}
	return
}
