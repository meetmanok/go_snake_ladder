package domain

import (
	"math/rand"
	"time"
)

type Dice struct {
}

func (d Dice) RollDice() int8 {
	rand.Seed(time.Now().UnixNano())
	return randomInt(1, 6)
}

func randomInt(min, max int8) int8 {
	return min + int8(rand.Intn(int(max-min)))
}

func NewDice() Dice {
	return Dice{}
}
