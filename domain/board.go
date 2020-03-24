package domain

type Board struct {
	Dice Dice
	Pawn Pawn
	Snakes Snakes
	Ladders Ladders
}

func NewBoard(dice Dice, pawn Pawn, snakes Snakes, ladders Ladders) Board {
	return Board{
		Dice: dice,
		Pawn:    pawn,
		Snakes:  snakes,
		Ladders: ladders,
	}
}