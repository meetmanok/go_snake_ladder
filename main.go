package main

import (
	"fmt"
	d "github.com/meetmanok/go_snake_ladder/domain"
	"os"
)

func main(){
	board := d.NewBoard(d.NewDice(), d.NewPawn(), d.NewSnakes(), d.NewLadders())

	for board.Pawn.GetPos() >= 0 && board.Pawn.GetPos() < 100 {
		var cRollNum int8
		var char string

		//Get user input to proceed for rolling dice
		fmt.Printf("Roll dice? (y/n): ")
		fmt.Scan(&char)
		switch char{
		case "y", "Y", "yes", "Yes", "YES":
			cRollNum = board.Dice.RollDice()
		case "n", "N", "no", "No", "NO":
			fmt.Printf("Quiting the game :-(")
			os.Exit(1)
		default:
			continue
		}

		//Move the pawn according to point got from dice rolling
		fmt.Printf("Current roll point: %d\n", cRollNum)
		if (board.Pawn.GetPos() + cRollNum) <= 100 {
			board.Pawn.SetPos(board.Pawn.GetPos() + cRollNum)
		}
		fmt.Printf("Pawn position: %d\n", board.Pawn.GetPos())

		//Verify that the current position of pawn is having Snake head or ladder foot
		//if current pawn pos is Snake's head pawn will be move to snake's tail
		//if current pawn pos is Ladder's foot pawn will be move to Ladder's top
		if v, ok := board.Snakes.GetSnakesPos(board.Pawn.GetPos()); ok {
			fmt.Printf("Oops, bitten by snake back to: %d\n", v)
			board.Pawn.SetPos(v)
		} else if v, ok := board.Ladders.GetLaddersPos(board.Pawn.GetPos()); ok {
			fmt.Printf("Hurray, got the ladder climb to: %d\n", v)
			board.Pawn.SetPos(v)
		}
	}
	if board.Pawn.GetPos() == 100 {
		fmt.Printf("Game Completed Successfully :-)")
	}
}
