package main

import (
	"P318UserSystem/server/model"
	process2 "P318UserSystem/server/process"
	"fmt"
	"net"
	"time"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 这里需要延时关闭
	defer conn.Close()

	processor := &process2.Processor{
		Conn: conn,
	}

	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误=", err)
		return
	}

	// 循环的客户端发送的信息
	// for {
	// buf := make([]byte, 8096)
	// fmt.Println("读取客户端发送的信息")
	// n, err := conn.Read(buf[0:4])
	// if n != 4 || err != nil {
	// 	fmt.Println("conn.Read err=", err)
	// 	return
	// }

	// mes, err := readPkg(conn)
	// if err != nil {
	// 	if err == io.EOF {
	// 		fmt.Println("客户端退出，服务器也退出...")
	// 		return
	// 	} else {
	// 		fmt.Println("readPkg err=", err)
	// 		return
	// 	}
	// }

	// err = serverProcessMes(conn, &mes)
	// if err != nil {
	// 	return
	// }

	// fmt.Println("mes=", mes)

	// fmt.Println("读到的buf =", buf[0:4])
	// }
}

// 这里我们编译一个函数，完成对 UserDao 的初始化任务
func initUserDao() {
	// 这里的 pool 本身就是一个全局变量
	model.MyUserDao = model.NewUserDao(model.Pool)
}

func main() {
	// 当服务器启动时，我们就去初始化我们的 redis 的链接池
	model.InitPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	// 提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	// 一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		// 一旦连接成功， 则启动一个协程和客户端保持通信。。。
		go process(conn)
	}
}
