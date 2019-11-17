package entity

type TableUser struct {
	UserId     int64  // 用户唯一ID
	Table      *Table // 用户所在桌
	SeatNo     int8   // 该用户的座位号
	Ready      bool   // 用户是否准备
	HandPokers Pokers // 用户剩余手牌
}
