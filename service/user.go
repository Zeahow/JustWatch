package service

import (
	"github.com/zeahow/just_watch/conf"
	"github.com/zeahow/just_watch/dal"
	"github.com/zeahow/just_watch/infrastructure/protocol"
	"github.com/zeahow/just_watch/infrastructure/transport"
	"sync"
)

type User int64

var (
	userMutex = sync.RWMutex{}
	userId2Client = make(map[int64]*Client, conf.DefaultClientNum)
)


type userService struct {
	sync.RWMutex
	userId2Client map[int64]*Client // ReqUserId => Client信息
}

/*
 * 检查用户是否已登录
 */
func (u User) IsLogged() bool {
	userMutex.RLock()
	defer userMutex.RUnlock()

	_, isLogged := userId2Client[int64(u)]
	return isLogged
}

/*
 * 检查连接t和用户id是否匹配
 */
func (u User) IsTransMatched(t transport.Transporter) bool {
	userMutex.RLock()
	defer userMutex.RUnlock()

	return u.IsLogged() && userId2Client[int64(u)].Transporter == t
}

/*
 * 登录
 */
func (u User) Login(t transport.Transporter, reqMsg *protocol.RequestMessage) error {
	resp := &protocol.ResponseMessage{MsgType: protocol.LoginResp}
	client := &Client{Transporter: t}

	userInfo, ok := reqMsg.OptionalBody.(protocol.UserInfo)
	if !ok {
		return client.Response(resp.WithRequired(0, false, "数据错误"))
	}

	userId := dal.UserDal.GetUserId(userInfo.UserName, userInfo.Password)
	if userId == 0 {
		return client.Response(resp.WithRequired(0, false, "账号或密码错误"))
	}

	userMutex.Lock()
	defer userMutex.Unlock()
	userId2Client[userId] = &Client{
		UserId:      userId,
		Transporter: t,
	}
	resp.ReqUserId = userId
	return client.Response(resp.WithRequired(0, true, ""))
}
