package process

import (
	"P318UserSystem/common/message"
	"P318UserSystem/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) {
	// 1. 创建一个 mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 3. 序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal failed =", err)
		return
	}
	mes.Data = string(data)

	// 4. 对 mes 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal failed =", err)
		return
	}

	// 5. 将 mes 发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	// 6. 发送
	err = tf.WritePkg(data)
}
