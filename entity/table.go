package entity

import (
	"sync"
	"time"
)

type Table struct {
	sync.RWMutex // 并发访问控制锁

	Timer          *time.Timer  // 超时计时器
	Id             int64        // 桌号
	Status         int8         // 该桌游戏状态
	Users          []*TableUser // 用户列表。User[i]表示该桌的第i位用户，User[i]为nil时，第i位座位无人。
	CreatorUser    *TableUser   // 创建用户Id
	TurnUser       *TableUser   // 当前出牌中用户Id
	LastShotUser   *TableUser   // 上次出牌用户Id
	LastShotPokers Pokers       // 上次出的牌
	RemainedPokers Pokers       // 该桌所剩余的牌
	ShotPokers     Pokers       // 该桌已打出的牌
	Multiple       int64        // 翻倍倍数
}
