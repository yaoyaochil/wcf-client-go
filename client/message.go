package client

import (
	"fmt"

	pb "github.com/yaoyaochil/wcf-client-go/proto"
	"google.golang.org/protobuf/proto"
)

// SendText 发送文本消息
func (c *ClientWCF) SendText(receiver, message string, aters string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_SEND_TXT,
		Msg: &pb.Request_Txt{
			Txt: &pb.TextMsg{
				Msg:      message,
				Receiver: receiver,
				Aters:    aters,
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("send text failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// SendImage 发送图片消息
func (c *ClientWCF) SendImage(receiver, imagePath string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_SEND_IMG,
		Msg: &pb.Request_File{
			File: &pb.PathMsg{
				Path:     imagePath,
				Receiver: receiver,
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("send image failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// SendFile 发送文件消息
func (c *ClientWCF) SendFile(receiver, filePath string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_SEND_FILE,
		Msg: &pb.Request_File{
			File: &pb.PathMsg{
				Path:     filePath,
				Receiver: receiver,
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("send file failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// SendXml 发送XML消息
func (c *ClientWCF) SendXml(receiver, content, path string, msgType uint64) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_SEND_XML,
		Msg: &pb.Request_Xml{
			Xml: &pb.XmlMsg{
				Receiver: receiver, // 接收者
				Content:  content,  // 内容
				Path:     path,     // 路径
				Type:     msgType,  // 消息类型
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("send xml failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// SendPatMsg 发送拍一拍消息
func (c *ClientWCF) SendPatMsg(roomid, wxid string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_SEND_PAT_MSG,
		Msg: &pb.Request_Pm{
			Pm: &pb.PatMsg{
				Roomid: roomid,
				Wxid:   wxid,
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("send pat message failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// ForwardMsg 转发消息
func (c *ClientWCF) ForwardMsg(msgid uint64, receiver string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_FORWARD_MSG,
		Msg: &pb.Request_Fm{
			Fm: &pb.ForwardMsg{
				Id:       msgid,
				Receiver: receiver,
			},
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("forward message failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// RevokeMsg 撤回消息
func (c *ClientWCF) RevokeMsg(msgid uint64) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_REVOKE_MSG,
		Msg: &pb.Request_Ui64{
			Ui64: msgid,
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

	if response.Msg == nil {
		return nil
	}

	switch msg := response.Msg.(type) {
	case *pb.Response_Status:
		if msg.Status != 0 {
			return fmt.Errorf("revoke message failed with status: %v", msg.Status)
		}
		return nil
	case *pb.Response_Str:
		return fmt.Errorf("received string response: %v", msg.Str)
	default:
		return fmt.Errorf("unexpected response type: %T", response.Msg)
	}
}

// 其他消息相关方法...
