package network

import (
	"net"
)

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	ClientIP() string
	ClientOrigin() string
	Close()
	Destroy()
}
