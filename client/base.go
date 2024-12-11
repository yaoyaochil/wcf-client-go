package client

import (
	"fmt"
	"time"

	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/pair1"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

type ClientWCF struct {
	socket mangos.Socket
}

// NewClient 创建一个新的 Base 客户端
func NewClient(addr string) (*ClientWCF, error) {
	socket, err := pair1.NewSocket()
	if err != nil {
		return nil, fmt.Errorf("failed to create socket: %v", err)
	}

	err = socket.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	err = socket.SetOption(mangos.OptionRecvDeadline, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to set recv timeout: %v", err)
	}

	return &ClientWCF{
		socket: socket,
	}, nil
}

// Close 关闭连接
func (c *ClientWCF) Close() {
	if c.socket != nil {
		c.socket.Close()
	}
}

// call 发送请求并接收响应
func (c *ClientWCF) call(req []byte) ([]byte, error) {
	err := c.socket.Send(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send: %v", err)
	}

	msg, err := c.socket.Recv()
	if err != nil {
		return nil, fmt.Errorf("failed to recv: %v", err)
	}

	return msg, nil
}
