package transport

import (
	"github.com/zeahow/just_watch/infrastructure/protocol"
)

/*
 * 连接处理函数
 */
type Handler func(transporter Transporter)

type Transporter interface {
	/*
	 * 初始化传输协议，并传入一个Handler，用于在接受请求后，对请求进行处理
	 */
	Init(handler Handler)

	/*
	 * 读取数据，返回一个Message与错误error
	 */
	Read() (*protocol.RequestMessage, error)

	/*
	 * 写入数据
	 */
	Write(msg *protocol.ResponseMessage) error
}

/*
 * 保存已注册Transporter的Name到Transporter实例的映射
 */
var Name2Transporter map[string]Transporter

/*
 * 注册一个Transporter
 */
func register(name string, transporter Transporter) {
	Name2Transporter[name] = transporter
}
