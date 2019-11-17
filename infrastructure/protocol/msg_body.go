package protocol

import "github.com/zeahow/just_watch/entity"

// 必须的Request信息
type RequiredReqBody struct {
	UserId string // 用户Id
}

// 必须的Response信息
type RequiredRespBody struct {
	MsgId   int64  // 消息Id，用于客户端判断有无丢失消息
	UserId  int64  // 消息请求用户Id，用于客户端判断更新谁的状态
	Success bool   // 请求是否成功
	Msg     string // 错误提示信息
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

type RegisterReqBody struct {
	UserName string
	Password string
}

type LoginReqBody struct {
	UserName string
	Password string
}

type NewTableReqBody struct {
	TableId int64
}

type JoinTableReqBody struct {
	TableId int64
}

type ExitTableReqBody struct {
	TableId int64
}

type DealPokersRespBody struct {
	Pokers entity.Pokers
}

type ShotPokersReqBody struct {
	Pokers entity.Pokers
}

type ShotPokersRespBody struct {
	Pokers entity.Pokers
}




type RegisterRespBody struct {
}

type LoginRespBody struct {
}

type LogoutReqBody struct {
}

type LogoutRespBody struct {
}

type NewTableRespBody struct {
}

type JoinTableRespBody struct {
}

type ExitTableRespBody struct {
}

type GetReadyReqBody struct {
}

type GetReadyRespBody struct {
}

type HangUpReqBody struct {
}

type HangUpRespBody struct {
}

type StartGameReqBody struct {
}

type StartGameRespBody struct {
}

type DealPokersReqBody struct {
}

type GameOverReqBody struct {
}

type GameOverRespBody struct {
}

type SyncTableInfoReqBody struct {
}

type SyncTableInfoRespBody struct {
}
