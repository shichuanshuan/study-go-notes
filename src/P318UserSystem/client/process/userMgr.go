package process

import (
	"P318UserSystem/client/model"
	"P318UserSystem/common/message"
	"fmt"
)

// 客户端要维护的 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	// 遍历一把 onlineUsers
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		// 不显示自己
		fmt.Println("用户id:\t", id)
	}
}

// 编写一个方法，处理返回的 notifyUserStatuseMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { // 原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
