package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Cell struct {
	Value     int  `json:"value"`
	IsMine    bool `json:"mine"`
	IsClicked bool `json:"clicked"`
	IsFlagged bool `json:"flagged"`
}

type Grid []Cell

type Game struct {
	Id           string  `json:"id"`
	Rows         int     `json:"rows"`
	Cols         int     `json:"cols"`
	Board        []Grid  `json:"board,omitempty"`
	Mines        int     `json:"mines"`
	NumClicks    int     `json:"numClicks,omitempty"`
	State        string  `json:"state,omitempty"`
	CreateTime   float64 `json:"createTime"`
	ConsumedTime float64 `json:"consumedTime"`
}

// Game are stored in this map for now
var games map[string]*Game

func init() {
	rand.Seed(time.Now().Unix())
	games = make(map[string]*Game)
}

func Create(g *Game) error {
	if g.Rows == 0 {
		return errors.New("Game need a number of rows")
	}
	if g.Cols == 0 {
		return errors.New("Game need a number of columns")
	}
	if g.Mines == 0 {
		return errors.New("Game need a number of mines")
	}

	g.Id = uuid.New().String()
	g.State = "created"
	games[g.Id] = g

	return nil
}

func Start(id string) (*Game, error) {
	var gameToStart *Game
	var ok bool
	if gameToStart, ok = games[id]; !ok {
		return nil, errors.New("Game not found")
	}

	gameToStart.CreateBoard()
	gameToStart.CreateTime = float64(time.Now().Unix())
	gameToStart.State = "started"
	return gameToStart, nil
}

func (g *Game) UpdateTimer() {
	g.ConsumedTime = float64(time.Now().Unix()) - g.CreateTime
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
			} else if cell.IsClicked {
				str += "c "
			} else {
				str += fmt.Sprintf("%d ", cell.Value)
			}
		}
		str += fmt.Sprintf("\n")
	}
	str += fmt.Sprintf("\n")
	return string(enc) + str, nil
}

func (g *Game) revealCell(row, col int) error {
	if g.Board[row][col].IsClicked {
		return errors.New("Cell already clicked")
	}
	g.Board[row][col].IsClicked = true

	if !g.Board[row][col].IsMine {
		g.NumClicks += 1
		g.revealAdjacents(row, col)
		if g.NumClicks == ((g.Rows * g.Cols) - g.Mines) {
			g.State = "won"
		}
	} else {
		g.State = "exploded"
	}

	return nil
}

func (g *Game) flagCell(row, col int) error {
	if g.Board[row][col].IsClicked {
		return errors.New("Cell already clicked")
	}
	g.Board[row][col].IsFlagged = true

	return nil
}

func (g *Game) revealAdjacents(row, col int) {

	// go down
	for i := row + 1; i < g.Rows; i++ {
		if g.Board[i][col].IsClicked || g.Board[i][col].IsMine || g.Board[i][col].Value != 0 {
			break
		}
		g.expandCell(i, col)
	}
	// go up
	for i := row - 1; i >= 0; i-- {
		if g.Board[i][col].IsClicked || g.Board[i][col].IsMine || g.Board[i][col].Value != 0 {
			break
		}
		g.expandCell(i, col)
	}
	// go left
	for i := col - 1; i >= 0; i-- {
		if g.Board[row][i].IsClicked || g.Board[row][i].IsMine || g.Board[row][i].Value != 0 {
			break
		}
		g.expandCell(row, i)
	}
	// go right
	for i := row + 1; i < g.Rows; i++ {
		if g.Board[row][i].IsClicked || g.Board[row][i].IsMine || g.Board[row][i].Value != 0 {
			break
		}
		g.expandCell(row, i)
	}

	// go down right
	for i := row + 1; i < g.Rows; i++ {
		if g.Board[i][i].IsClicked || g.Board[i][i].IsMine || g.Board[i][i].Value != 0 {
			break
		}
		g.expandCell(i, i)
	}
	// go down left
	j := (col - 1)
	for i := (row + 1); i < g.Rows && j >= 0; i++ {
		if g.Board[i][j].IsClicked || g.Board[i][j].IsMine || g.Board[i][j].Value != 0 {
			break
		}
		g.expandCell(i, j)
		j--
	}

	// go up right
	j = (col + 1)
	for i := row - 1; i >= 0 && j < g.Cols; i-- {
		if g.Board[i][j].IsClicked || g.Board[i][j].IsMine || g.Board[i][j].Value != 0 {
			break
		}
		g.expandCell(i, j)
		j++
	}

	// go up left
	for i := row - 1; i >= 0; i-- {
		if g.Board[i][i].IsClicked || g.Board[i][i].IsMine || g.Board[i][i].Value != 0 {
			break
		}
		g.expandCell(i, i)
	}
}

func (g *Game) expandCell(row, col int) {
	if g.Board[row][col].Value == 0 && !g.Board[row][col].IsMine {
		g.Board[row][col].IsClicked = true
		g.NumClicks++
	}
}
