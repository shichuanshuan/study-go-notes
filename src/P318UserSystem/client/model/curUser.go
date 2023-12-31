package model

import (
	"P318UserSystem/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
