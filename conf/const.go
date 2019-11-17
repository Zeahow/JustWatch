package conf

import "time"

const (
	DefaultClientNum  = 1024             // 默认客户端数量
	WaiteTime         = 30 * time.Second // 默认超时时间
	DefaultTableNum   = 1024             // 默认桌数
	TableMaxUserNum   = 6                // 一桌最大人数
	TableMinUserNum   = 4                // 一桌最小人数
	RetryOnFailedTime = 3                // 失败重试次数
)
