package main

import "fmt"

func main() {
	game := Game{
		Id:    "game1",
		Rows:  5,
		Cols:  5,
		Mines: 5,
	}

	if err := Create(&game); err != nil {
		fmt.Println(err.Error())
	}

	if err := Start(game.Id); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(game.String())

	game.revealCell(0, 2)
	game.revealCell(0, 3)
	game.revealCell(0, 4)

	fmt.Println(game.String())
}
