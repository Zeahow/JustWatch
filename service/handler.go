package service

import (
	"fmt"
	"github.com/zeahow/just_watch/infrastructure/protocol"
	"github.com/zeahow/just_watch/infrastructure/transport"
)

func Handle(t transport.Transporter) {
	var reqMsg *protocol.RequestMessage
	for err := error(nil); err != nil; logErr(err) {
		reqMsg, err = t.Read()
		if reqMsg == nil || reqMsg.MsgType == protocol.UnknownMsg {
			continue
		}
		user := User(reqMsg.UserId)

		switch reqMsg.MsgType {
		case protocol.LoginReq: // 登录
			err = user.Login(t, reqMsg)
			continue
		case protocol.RegisterReq:
			// todo 注册
			continue
		default: // 其它的消息类型需校验登录状态
			if !user.IsTransMatched(t) {continue}
		}

		switch reqMsg.MsgType {
		case protocol.LogoutReq:

		}
	}

}

func logErr(err error) {
	if err !=nil {
		fmt.Println(err)
	}
}
