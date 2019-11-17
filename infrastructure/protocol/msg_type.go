package protocol

/*
 * 消息类型。默认零值为未知类型。
 * Req为client => server的请求。
 * Resp为server => client的回复。
 */
const (
	UnknownMsg    int32 = iota // 未知
	RegisterReq                // 注册Request
	RegisterResp               // 注册Response
	LoginReq                   // 登录Request
	LoginResp                  // 登录Response
	LogoutReq                  // 注销Request
	LogoutResp                 // 注销Response
	NewTableReq                // 创建游戏桌Request
	NewTableResp               // 创建游戏桌Response
	JoinTableReq               // 加入游戏桌Request
	JoinTableResp              // 加入游戏桌Response
	ExitTableReq               // 退出游戏桌Request
	ExitTableResp              // 退出游戏桌Response
	GetReadyReq                // 准备Request
	GetReadyResp               // 准备Response
	HangUpReq                  // 取消准备Request
	HangUpResp                 // 取消准备Response
	StartGameReq               // 开始游戏Request
	StartGameResp              // 开始游戏Response
	DealPokersReq              // 发牌Request
	DealPokersResp             // 发牌Response
	ShotPokersReq              // 出牌Request
	ShotPokersResp             // 出牌Response
	GameOverReq                // 游戏结束Request
	GameOverResp               // 游戏结束Response
	SyncTableInfoReq           // 同步游戏桌信息Request
	SyncTableInfoResp          // 同步游戏桌信息Response
)
