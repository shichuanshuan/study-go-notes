package process2

import (
	"P318UserSystem/common/message"
	"P318UserSystem/server/utils"
	"fmt"
	"io"
	"net"
)

// 先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

// 编写一个函数 serverProcessMes 函数
// 功能： 根据客户端发送消息种类不通，决定调用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	fmt.Println("accept message is", mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		// 创建一个UserProcess 实例
		up := &UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 群发消息
		smsprocess := SmsProcess{}
		smsprocess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}

	return
}

func (this *Processor) Process2() (err error) {

	// 循环的客户端发送的信息
	for {
		// 这里我们将读取数据包，直接封装成一个函数readPkg(), 返回message, Err
		// 创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}

		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
