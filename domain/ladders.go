package domain

type Ladders struct {
	laddersPos map[int8]int8
}

func NewLadders() Ladders {
	return Ladders{laddersPos:map[int8]int8{1:38, 4:14, 8:10, 21:42, 28:76, 50:67, 71:92, 88:99}}
}

func (l Ladders) GetLaddersPos(startPos int8) (int8, bool) {
	v, ok := l.laddersPos[startPos]
	return v, ok
}