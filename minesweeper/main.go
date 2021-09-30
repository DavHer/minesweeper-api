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
}
