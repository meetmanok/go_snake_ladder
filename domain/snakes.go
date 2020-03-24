package domain

type Snakes struct {
	snakesPos map[int8]int8
}

func NewSnakes() Snakes {
	return Snakes{snakesPos:map[int8]int8{36:6, 32:10, 62:18, 88:24, 48:26, 95:56, 97:78}}
}

func (s Snakes) GetSnakesPos(startPos int8) (int8, bool) {
	v, ok := s.snakesPos[startPos]
	return v, ok
}