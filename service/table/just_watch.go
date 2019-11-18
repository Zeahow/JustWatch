package table

import (
	"github.com/zeahow/just_watch/conf"
	"github.com/zeahow/just_watch/entity"
	"github.com/zeahow/just_watch/rule"
	"github.com/zeahow/just_watch/service"
	"math/rand"
	"sync"
	"time"
)

type Table entity.Table

const (
	GameWaitingStart = iota // 等待游戏开始
	GameRunning             // 游戏进行中
	GameOver                // 游戏结束
)

var (
	curTableId int64 = 10000
	tableMutex       = sync.RWMutex{}

	tableId2Table = make(map[int64]*Table, conf.DefaultTableNum)
)

func TableOf(tableId int64) *Table {
	tableMutex.RLock()
	defer tableMutex.RUnlock()
	return tableId2Table[tableId]
}

func NewTable(creatorUserId int64) *Table {
	tableMutex.Lock()
	defer tableMutex.Unlock()
	curTableId += rand.Int63n(100)
	user := &entity.TableUser{
		UserId: creatorUserId,
		SeatNo: 0,
	}
	table := &Table{
		Id:          curTableId,
		Users:       make([]*entity.TableUser, conf.TableMaxUserNum),
		CreatorUser: user,
	}
	user.Table = (*entity.Table)(table)
	table.Users[0] = user
	table.Reset()
	tableId2Table[table.Id] = table
	return table
}

/*
 * 重置桌的状态。包括重置剩余牌、已出牌、当前出牌用户等。
 */
func (t *Table) Reset() {
	t.Lock()
	defer t.Unlock()
	t.RemainedPokers, t.ShotPokers = entity.New54Pokers(), entity.Pokers{}
	if t.LastShotUser != nil {
		t.TurnUser = t.LastShotUser
	} else {
		t.TurnUser = t.FindNextUser(-1)
	}
	t.LastShotUser, t.LastShotPokers = nil, nil
	t.Status = GameWaitingStart
	t.Multiple = 1
	t.Timer = nil
}

/*
 * 查找curSeatNo的下一个用户
 */
func (t *Table) FindNextUser(curSeatNo int) *entity.TableUser {
	// 先从curSeatNo往后找
	for i := curSeatNo + 1; curSeatNo < len(t.Users); i++ {
		if t.Users[i] != nil {
			return t.Users[i]
		}
	}
	// 找不到再从头找，直到curSeatNo
	for i := 0; i <= curSeatNo; i++ {
		if t.Users[i] != nil {
			return t.Users[i]
		}
	}
	return nil
}

/*
 * 查找可用的座位号
 */
func (t *Table) findAvailableSeatNo() int8 {
	t.RLock()
	defer t.RUnlock()

	for seatNo, user := range t.Users {
		if user == nil {
			return int8(seatNo)
		}
	}
	return -1
}

/*
 * 新加一个用户
 */
func (t *Table) AddUser(userId int64) (success bool) {
	t.Lock()
	defer t.Unlock()

	seatNo := t.findAvailableSeatNo()
	if seatNo == -1 {
		return false
	}

	newUser := &entity.TableUser{
		UserId: userId,
		Table:  (*entity.Table)(t),
		SeatNo: seatNo,
	}
	t.Users[seatNo] = newUser
	return true
}

/*
 * 用户退出。如果是游戏中状态，将该用户托管给机器人。
 */
func (t *Table) ExitUser(userId int64) {
	t.Lock()
	defer t.Unlock()

	for i, user := range t.Users {
		if user.UserId == userId {
			if t.Status == GameRunning {    // 游戏进行中，则该用户退出后，为该用户设置一个机器人
				// todo 设置机器人
			} else {
				t.Users[i] = nil
			}
		}
	}
}

/*
 * 开始游戏
 */
func (t *Table) StartGame() bool {
	t.Lock()
	defer t.Unlock()

	// 判断是否所有玩家都已准备，并且在座玩家大于4
	inSeatUserNum := 0
	for _, user := range t.Users {
		if user != nil && user.Ready == false {
			return false
		} else if user != nil && user.Ready == true {
			inSeatUserNum++
		}
	}
	if inSeatUserNum < conf.TableMinUserNum {
		return false
	}

	t.Reset()
	t.Status = GameRunning
	t.Timer = time.AfterFunc(conf.WaiteTime, t.onTimeOut)   // 开始出牌倒计时
	return true
}

/*
 * 出牌超时，自动不出牌
 */
func (t *Table) onTimeOut() {
	(*service.TableUser)(t.TurnUser).ShotPokers(entity.Pokers{})
}

/*
 * 用户出牌
 * param userId: 要出牌的用户
 * param shotPokers: 要出的牌
 * return: 出牌是否成功
 */
func (t *Table) OnUserShot(userId int64, shotPoker entity.Pokers) bool {
	turnUser := t.TurnUser
	if userId != turnUser.UserId {
		return false
	}
	if canShot, remainedPokers := rule.JustWatchRule.CanShot(t.LastShotPokers, turnUser.HandPokers, shotPoker); canShot {
		turnUser.HandPokers = remainedPokers
		t.TurnUser = t.FindNextUser(int(turnUser.SeatNo))
		t.LastShotPokers = shotPoker
		t.LastShotUser = turnUser
		// 同步出牌给所有用户
		t.Sync2AllUsers()
	} else {
		return false
	}
	return true
}

/*
 * 同步至所有用户
 */
func (t *Table) Sync2AllUsers() {
	for _, user := range t.Users {
		if user != nil {
			t.sync(user)
		}
	}
}

func (t *Table) sync(user *entity.TableUser) {

}
