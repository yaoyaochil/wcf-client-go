package client

import (
	"fmt"

	pb "github.com/yaoyaochil/wcf-client-go/proto"
	"google.golang.org/protobuf/proto"
)

// IsLogin 检查是否已登录
func (c *ClientWCF) IsLogin() (bool, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_IS_LOGIN,
		Msg: &pb.Request_Empty{
			Empty: &pb.Empty{},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return false, err
	}

	resp, err := c.call(data)
	if err != nil {
		return false, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return false, err
	}

	// 如果响应消息为空，认为是已登录状态
	if response.Msg == nil {
		return true, nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		return msg.Status == 1, nil
	case *pb.Response_Str:
		return false, fmt.Errorf("received string response: %v", msg.Str)
	default:
		return false, fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// RefreshQrcode 刷新登录二维码
func (c *ClientWCF) RefreshQrcode() error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_REFRESH_QRCODE,
		Msg: &pb.Request_Empty{
			Empty: &pb.Empty{},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := c.call(data)
	if err != nil {
		return err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return err
	}

	status, ok := response.Msg.(*pb.Response_Status)
	if !ok || status.Status != 0 {
		return fmt.Errorf("refresh qrcode failed with status: %v", status.Status)
	}

	return nil
}
