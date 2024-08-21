package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	ClientIP() string
	IsWebSocket() (bool,interface{})
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
