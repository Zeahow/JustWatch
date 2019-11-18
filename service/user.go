package service

import (
	"github.com/zeahow/just_watch/dal"
	"github.com/zeahow/just_watch/infrastructure/protocol"
	"github.com/zeahow/just_watch/infrastructure/transport"
)

type User int64

func UserOf(userId int64) User {
	return User(userId)
}

/*
 * 检查用户是否已登录
 */
func (u User) IsLogged() bool {
	return ClientOf(int64(u)) != nil
}

/*
 * 检查连接t和用户id是否匹配
 */
func (u User) IsTransMatched(t transport.Transporter) bool {
	client := ClientOf(int64(u))
	return client != nil && client.Transporter == t
}

/*
 * 注册
 */
func (u User) Register(t transport.Transporter, reqMsg *protocol.RequestMessage) (userId int64, err error) {
	return 0, nil
}

/*
 * 检查账号密码是否合法
 */
func (u User) Login(reqMsg *protocol.RequestMessage, respMsg *protocol.ResponseMessage) *protocol.ResponseMessage {
	userInfo, ok := reqMsg.OptionalBody.(protocol.UserInfo)
	respMsg = &protocol.ResponseMessage{MsgType: protocol.LoginResp}
	if !ok {
		return respMsg.WithRequired(0, false, protocol.WrongRequest)
	}

	userId := dal.UserDal.GetUserId(userInfo.UserName, userInfo.Password)
	if userId == 0 {
		return respMsg.WithRequired(0, false, protocol.WrongUserInfo)
	}
	reqMsg.UserId = userId  // 回写登录后的UserId到请求消息里
	return respMsg.WithRequired(userId, true, protocol.NoErrors)
}
