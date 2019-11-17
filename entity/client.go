package entity

import "github.com/zeahow/just_watch/infrastructure/transport"

type Client struct {
	UserId      int64
	MsgId       int64   // 消息Id，用于客户端判断有无丢失消息
	Transporter transport.Transporter
}
