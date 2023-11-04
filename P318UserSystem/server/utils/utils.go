package utils

import (
	"P318UserSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte // 这时传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据...")
	// conn.Read 在conn没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了 conn, 则就不会阻塞
	n1, err := this.Conn.Read(this.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// 接收数据的长度
	fmt.Println("frist read =", this.Buf[:n1])

	// 根据buf[:4] 转成一个 uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:])

	// 根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	// fmt.Println("pkglen =", int(pkgLen))
	// fmt.Println("n =", n)
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	// 把 pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// 发送 data 本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	return
}
