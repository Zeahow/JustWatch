package service

import (
	"fmt"
	"github.com/zeahow/just_watch/infrastructure/protocol"
	"github.com/zeahow/just_watch/infrastructure/transport"
)

func Handle(t transport.Transporter) {
	var reqMsg *protocol.RequestMessage
	var respMsg = &protocol.ResponseMessage{}
	var client *Client
	for err := error(nil); err != nil; logErr(err) {
		reqMsg, err = t.Read()
		respMsg.MsgType = protocol.CommonResp

		if client = getClient(t, reqMsg); reqMsg == nil || client == nil {
			respMsg = respMsg.WithRequired(0, false, protocol.WrongRequest)
			continue
		}

		userId := respMsg.ReqUserId
		switch reqMsg.MsgType {
		case protocol.LoginReq: // 登录
			respMsg.MsgType = protocol.LoginResp
			if respMsg = UserOf(userId).Login(reqMsg, respMsg); respMsg.Success {
				client.Save()
			}
		case protocol.RegisterReq:
			// todo 注册
			continue
		default: // 其它的消息类型需校验登录状态
			respMsg = respMsg.WithRequired(0, false, protocol.WrongRequest)
		}
	}

}

func getClient(t transport.Transporter, reqMsg *protocol.RequestMessage) (client *Client) {
	if reqMsg == nil || reqMsg.MsgType == protocol.UnknownMsg {
		return nil
	}

	client = ClientOf(reqMsg.UserId)
	// 登录或注册需生成临时Client
	if reqMsg.MsgType == protocol.LoginResp || reqMsg.MsgType == protocol.RegisterReq {
		client = TempClient(t)
	}

	if client == nil || client.Transporter != t {
		return nil
	}
	return client
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
