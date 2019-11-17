package service

import (
	"errors"
	"github.com/zeahow/just_watch/conf"
	"github.com/zeahow/just_watch/entity"
	"github.com/zeahow/just_watch/infrastructure/protocol"
)

type Client entity.Client

/*
 * 应答消息 msg 至 c.Transporter。再多失败重试 conf.RetryOnFailedTime 次。
 */
func (c *Client) Response(msg *protocol.ResponseMessage) error {
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
