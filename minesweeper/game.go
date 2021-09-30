package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	Value     int  `json:"value"`
	IsMine    bool `json:"mine"`
	IsClicked bool `json:"clicked"`
	IsFlagged bool `json:"flagged"`
}

type Grid []Cell

type Game struct {
	Id        string `json:"id"`
	Rows      int    `json:"rows"`
	Cols      int    `json:"cols"`
	Board     []Grid `json:"board"`
	Mines     int    `json:"mines"`
	NumClicks int    `json:"numClicks"`
	State     string `json:"state"`
}

// Game are stored in this map for now
var games map[string]*Game

func init() {
	rand.Seed(time.Now().Unix())
	games = make(map[string]*Game)
}

func Create(g *Game) error {
	if g.Id == "" {
		return errors.New("Game need an Id")
	}
	if g.Rows == 0 {
		return errors.New("Game need a number of rows")
	}
	if g.Cols == 0 {
		return errors.New("Game need a number of columns")
	}
	if g.Mines == 0 {
		return errors.New("Game need a number of mines")
	}

	g.State = "created"
	games[g.Id] = g

	return nil
}

func Start(id string) error {
	var gameToStart *Game
	var ok bool
	if gameToStart, ok = games[id]; !ok {
		return errors.New("Game not found")
	}

	gameToStart.CreateBoard()
	gameToStart.State = "started"
	return nil
}

func (g *Game) CreateBoard() error {
	g.Board = make([]Grid, g.Rows)
	for row := range g.Board {
		g.Board[row] = make(Grid, g.Cols)
	}

	// Set mines and values
	for i := 0; i < g.Mines; i++ {
		randRow := rand.Intn(g.Rows)
		randCol := rand.Intn(g.Cols)
		if !g.Board[randRow][randCol].IsMine {
			g.Board[randRow][randCol].IsMine = true
		} else {
			i--
		}

	}

	for i, row := range g.Board {
		for j, cell := range row {
			if cell.IsMine {
				setValuesForMine(g, i, j)
			}
		}
	}

	return nil
}

func setValuesForMine(game *Game, row, col int) {
	for i := row - 1; i < row+2; i++ {
		if i >= 0 && i < game.Rows {
			for j := col - 1; j < col+2; j++ {
				if j >= 0 && j < game.Cols && i != row && j != col {
					game.Board[i][j].Value++
				}
			}
		}
	}
}

func (g *Game) String() (string, error) {
	enc, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var str string = "\n"
	for _, row := range g.Board {
		for _, cell := range row {
			if cell.IsMine {
				str += "m "
			} else if cell.IsFlagged {
				str += "f "
			} else {
				str += fmt.Sprintf("%d ", cell.Value)
			}
		}
		str += fmt.Sprintf("\n")
	}
	str += fmt.Sprintf("\n")
	return string(enc) + str, nil
}

/*
str := fmt.Sprintf("Id: %b\n", g.Id)
	str += fmt.Sprintf("Rows: %b\n", g.Rows)
	str += fmt.Sprintf("Cols: %b\n", g.Cols)
	str += fmt.Sprintf("Mines: %b\n", g.Mines)
	str += fmt.Sprintf("NumClicks: %b\n", g.NumClicks)
	str += fmt.Sprintf("State: %b\n", g.State)
	str += fmt.Sprintf("Board: \n")

	return fmt.Sprintf("%b", g)
*/