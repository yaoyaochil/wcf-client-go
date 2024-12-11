package client

import (
	"fmt"

	pb "github.com/yaoyaochil/wcf-client-go/proto"
	"google.golang.org/protobuf/proto"
)

// AddRoomMembers 添加群成员
func (c *ClientWCF) AddRoomMembers(roomid string, wxids string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_ADD_ROOM_MEMBERS,
		Msg: &pb.Request_M{
			M: &pb.MemberMgmt{
				Roomid: roomid,
				Wxids:  wxids,
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

	status, ok := response.Msg.(*pb.Response_Status)
	if !ok || status.Status != 0 {
		return fmt.Errorf("add room members failed with status: %v", status.Status)
	}

	return nil
}

// DelRoomMembers 删除群成员
func (c *ClientWCF) DelRoomMembers(roomid string, wxids string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_DEL_ROOM_MEMBERS,
		Msg: &pb.Request_M{
			M: &pb.MemberMgmt{
				Roomid: roomid,
				Wxids:  wxids,
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

	status, ok := response.Msg.(*pb.Response_Status)
	if !ok || status.Status != 0 {
		return fmt.Errorf("delete room members failed with status: %v", status.Status)
	}

	return nil
}

// InvRoomMembers 邀请群成员
func (c *ClientWCF) InvRoomMembers(roomid string, wxids string) error {
	req := &pb.Request{
		Func: pb.Functions_FUNC_INV_ROOM_MEMBERS,
		Msg: &pb.Request_M{
			M: &pb.MemberMgmt{
				Roomid: roomid,
				Wxids:  wxids,
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

	status, ok := response.Msg.(*pb.Response_Status)
	if !ok || status.Status != 0 {
		return fmt.Errorf("invite room members failed with status: %v", status.Status)
	}

	return nil
}
