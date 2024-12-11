package client

import (
	"fmt"

	pb "github.com/yaoyaochil/wcf-client-go/proto"
	"google.golang.org/protobuf/proto"
)

// GetSelfWxid 获取自己的微信ID
func (c *ClientWCF) GetSelfWxid() (string, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_SELF_WXID,
		Msg: &pb.Request_Empty{
			Empty: &pb.Empty{},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := c.call(data)
	if err != nil {
		return "", err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return "", err
	}

	str, ok := response.Msg.(*pb.Response_Str)
	if !ok {
		return "", fmt.Errorf("unexpected response type")
	}

	return str.Str, nil
}

// GetContacts 获取通讯录
func (c *ClientWCF) GetContacts() ([]*pb.RpcContact, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_CONTACTS,
		Msg: &pb.Request_Empty{
			Empty: &pb.Empty{},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	contacts, ok := response.Msg.(*pb.Response_Contacts)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}

	return contacts.Contacts.Contacts, nil
}

// GetContactInfo 获取联系人信息
func (c *ClientWCF) GetContactInfo(wxid string) (*pb.RpcContacts, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_CONTACT_INFO,
		Msg: &pb.Request_Str{
			Str: wxid,
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.GetContacts(), nil
}

// GetSelfContactInfo 获取个人信息 Functions_FUNC_GET_USER_INFO
func (c *ClientWCF) GetSelfContactInfo() (*pb.UserInfo, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_USER_INFO,
		Msg:  nil,
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.GetUi(), nil
}
