package main

import (
	"P318UserSystem/client/process"
	"fmt"
	"os"
)

// 定义两个变量
var userId int
var userPwd string

func main() {
	// 接收用户的选择
	var key int

	// 判断是否还继续显示菜单
	// var loop = true

	for {
		fmt.Println("------------- 欢迎登陆多人聊天系统 -------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			// 1. 创建一个User Process 的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 2. 调用 UserProcess, 完成注册的请求、
			up := process.UserProcess{}
			up.Register(userId, userPwd)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重写输入")
		}
	}

}
