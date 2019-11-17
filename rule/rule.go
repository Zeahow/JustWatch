package rule

import (
	"github.com/zeahow/just_watch/entity"
)

type Rule interface {
	/*
	 * 比较两个扑克牌的大小
	 * return isLarger: -1→A<B; 0→A=B; 1→A>B
	 * return isDouble: 是否需要翻倍
	 */
	Compare(pokerA, pokerB entity.Pokers) (isLarger int8, isDouble bool)

	/*
	 * 判断能否出牌
	 * param lastUserShotPokers: 上一个用户出的牌
	 * param handPokers: 用户手中剩余的牌
	 * param wantToShotPokers: 本次想出的牌
	 * return canShot: 能否出牌
	 * return remainedPokers: 如果能出牌，出牌后手中剩余的牌
	 */
	CanShot(lastUserShotPokers entity.Pokers, handPokers, wantToShotPokers entity.Pokers) (
		canShot bool, remainedPokers entity.Pokers)

	EndGame()
}