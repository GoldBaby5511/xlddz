package network

import (
	"net"
)

type Conn interface {
	ReadMsg() (BaseMessage, []byte, error)
	WriteMsg(bm BaseMessage, msgData, otherData []byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
