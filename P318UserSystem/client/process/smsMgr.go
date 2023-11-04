package process

import (
	"P318UserSystem/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var sms message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &sms)
	if err != nil {
		fmt.Println("接收消息失败 err", err)
		return
	}

	fmt.Printf("用户id:%d 向大家发送的消息为:%s", sms.UserId, sms.Content)
	fmt.Println()
}
