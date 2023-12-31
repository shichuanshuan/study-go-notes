package main

import (
	"fmt"
	"net"
)

// 1.在计算机（尤其是做服务器）要尽可能的少开端口
// 2. 一个端口只能被一个程序监听
// 3.netstat -an	// 可以查看本机有哪些端口在监听
// 4.netstat -anb	// 来查看监听端口的pid，在结合任务管理器关闭不安全的端口

// 测试服务器方法
// 1. 写一个客户端程序
// 2. telnet + ip地址， 退出telnet ctrl + ]

func process(conn net.Conn) {
	// 这里我们循环的接收客户端发送的数据
	defer conn.Close() // 关闭conn

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1. 等待客户端通过conn发送信息
		// 2. 如果客户端没有writer
		fmt.Printf("服务器在等待客户端 %s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return
		}
		// 3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
		fmt.Println()
	}
}

func main() {
	fmt.Println("服务器开始监听")
	// net.Listen("tcp", "0.0.0.0:8888")
	// 1. tcp 表示使用网络协议是tcp
	// 2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close() // 延时关闭 listen

	// 循环等待客户端来链接我
	for {
		// 等待客户端
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() sucess con=%v\n", conn)
		}
		fmt.Println("==================================")
		go process(conn)
	}

}
