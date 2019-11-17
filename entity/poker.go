package entity

type Pokers []int

func New54Pokers() Pokers {
	pokers := make([]int, 0, 54)
	for i := 0; i < 54; i++ {
		pokers = append(pokers, i)
	}
	return pokers
}
