package protocol

import "github.com/zeahow/just_watch/entity"

// 必须的Request信息
type RequiredReqBody struct {
	UserId int64 // 用户Id
}

// 必须的Response信息
type RequiredRespBody struct {
	MsgId     int64  // 消息Id，用于客户端判断有无丢失消息
	ReqUserId int64  // 发起请求的用户的UserId，用于客户端判断更新谁的状态
	Success   bool   // 请求是否成功
	Msg       string // 错误提示信息
}

type UserInfo struct {
	UserName string
	Password string
}

type TableInfo struct {
	TableId int64
}

type PokersInfo struct {
	Pokers entity.Pokers
}
