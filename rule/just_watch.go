package rule

import "github.com/zeahow/just_watch/entity"

type JustWatchRule struct {

}

func (j JustWatchRule) Compare(pokerA, pokerB entity.Pokers) (isLarger int8, isDouble bool) {
	panic("implement me")
}

func (j JustWatchRule) CanShot(lastUserShotPokers entity.Pokers, handPokers, wantToShotPokers entity.Pokers) (
	canShot bool, remainedPokers entity.Pokers) {
	panic("implement me")
}

func (j JustWatchRule) EndGame() {
	panic("implement me")
}
