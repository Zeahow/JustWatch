package service

import (
	"errors"
	"github.com/zeahow/just_watch/conf"
	"github.com/zeahow/just_watch/entity"
	"github.com/zeahow/just_watch/infrastructure/protocol"
	"github.com/zeahow/just_watch/infrastructure/transport"
	"log"
	"sync"
)

type Client entity.Client

var (
	clientMutex = sync.RWMutex{}
	userId2Client = make(map[int64]*Client, conf.DefaultClientNum)  // UserId => Client信息
)

/*
 * 获取userId对应的Client
 */
func ClientOf(userId int64) *Client {
	clientMutex.RLock()
	defer clientMutex.RUnlock()
	return userId2Client[userId]
}

/*
 * 获取临时的client
 */
func TempClient(t transport.Transporter) *Client {
	return &Client{Transporter: t}
}

/*
 * 应答消息 msg 至 c.Transporter。再多失败重试 conf.RetryOnFailedTime 次。
 */
func (c *Client) Response(msg *protocol.ResponseMessage) (err error) {
	defer log.Printf("[Cilent.Reponse] client[%v] msg[%v] err[%v]", c, msg, err)

	if c == nil {
		return errors.New("client is nil")
	}
	msg.MsgId = c.MsgId
	for retryTime := 0; retryTime <= conf.RetryOnFailedTime; retryTime++ {
		if c.Transporter.Write(msg) == nil {
			c.MsgId++   // 发送成功，消息号+1
			return nil
		}
	}
	return errors.New("failed to response after try 3 times")
}

/*
 * 保存客户端
 */
func (c *Client) Save() {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	userId := c.UserId
	userId2Client[userId] = c
}
