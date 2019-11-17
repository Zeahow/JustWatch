package transport

import "github.com/zeahow/just_watch/infrastructure/protocol"

type WebSocket struct {

}

func init() {
	register("WebSocket", WebSocket{})
}

func (w WebSocket) Init(handle Handler) {
	go handle(w)
}

func (WebSocket) Read() (msg *protocol.Message, err error) {
	return
}

func (WebSocket) Write(msg *protocol.Message) (err error) {
	return
}