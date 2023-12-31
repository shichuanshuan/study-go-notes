package process2

import (
	"P318UserSystem/common/message"
	"P318UserSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

// 写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	// 遍历服务器端的 onlineUsers map[int]*UserProcess
	// 将消息转发去取出

	// 取出 mes 的内容 SmsMes
	var smsMes message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发数据失败 err=", err)
		return
	}
}
