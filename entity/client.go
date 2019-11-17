package entity

import "github.com/zeahow/just_watch/infrastructure/transport"

type Client struct {
	Uid         int64
	transporter transport.Transporter
}
