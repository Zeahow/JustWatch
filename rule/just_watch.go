package rule

import "github.com/zeahow/just_watch/entity"

var JustWatchRule = justWatchRule{}

type justWatchRule struct { }

func (j justWatchRule) Compare(pokerA, pokerB entity.Pokers) (isLarger int8, isDouble bool) {
	panic("implement me")
}

func (j justWatchRule) CanShot(lastUserShotPokers entity.Pokers, handPokers, wantToShotPokers entity.Pokers) (
	canShot bool, remainedPokers entity.Pokers) {
	panic("implement me")
}

func (j justWatchRule) EndGame() {
	panic("implement me")
}
