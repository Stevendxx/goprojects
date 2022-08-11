package main

import (
	"fmt"
	"net"
)

// 服务端

func main() {
	// 开启端口监听
	network := "tcp"
	address := "localhost:8991"
	fmt.Println("[服务器]:开启8991端口监听...")
	listener, errListen := net.Listen(network, address)
	if errListen != nil {
		fmt.Println("[错误]:开启监听错误,请检查网络方式和主机IP及端口.")
		fmt.Println(errListen)
		return
	}
	defer listener.Close()

	// 接受客户端请求并进行处理
	for {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			fmt.Println("[错误]:接受请求错误.")
			fmt.Println(errAccept)
			continue
		}
		// 开协程与客户端交互
		go handleConn(conn)
	}
}
