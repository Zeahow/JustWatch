package service

import (
	"github.com/zeahow/just_watch/conf"
	"github.com/zeahow/just_watch/entity"
	"sync"
)

type TableUser entity.TableUser

var (
	tableUserMutex   = sync.RWMutex{}
	userId2TableUser = make(map[int64]*TableUser, conf.DefaultTableNum)
)

/*
 * 从userId获取TableUser
 */
func TableUserOf(userId int64) *TableUser {
	tableUserMutex.RLock()
	defer tableUserMutex.RUnlock()
	return userId2TableUser[userId]
}

/*
 * 同步用户t所在游戏桌的信息给用户t
 */
func (t *TableUser) SyncTableInfo() error {
	if t == nil {
		return nil
	}
	// todo 同步table
	return nil
}

/*
 * 用户出牌
 */
func (t *TableUser) ShotPokers(pokers entity.Pokers) {

}
