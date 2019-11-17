package main

import (
	"github.com/zeahow/just_watch/infrastructure/transport"
	"github.com/zeahow/just_watch/service"
)



func main()  {
	for name, transporter := range transport.Name2Transporter {
		if name == "WebSocket" {
			transporter.Init(service.Handle)
		}
	}
}
