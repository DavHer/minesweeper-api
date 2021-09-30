package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardPos struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func StartApi() {
	router := gin.Default()
	router.POST("/minesweeper", createGame)
	router.POST("/minesweeper/:id/start", startGame)
	router.GET("/minesweeper/:id", getGame)
	router.POST("/minesweeper/:id/reveal", revealCell)
	router.POST("/minesweeper/:id/flag", flagCell)

	router.Run("localhost:8080")
}

func createGame(c *gin.Context) {
	var game Game
	if err := c.BindJSON(&game); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := Create(&game); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, game)
}

func getGame(c *gin.Context) {
	id := c.Param("id")
	var game *Game
	var ok bool
	if game, ok = games[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, game)
}

func startGame(c *gin.Context) {
	var game *Game
	var err error
	id := c.Param("id")
	if game, err = Start(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, game)
}

func revealCell(c *gin.Context) {
	var game *Game
	var err error
	var ok bool
	id := c.Param("id")
	var pos BoardPos

	if game, ok = games[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	if err = c.BindJSON(&pos); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = game.revealCell(pos.Row, pos.Col); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

func flagCell(c *gin.Context) {
	var game *Game
	var err error
	var ok bool
	id := c.Param("id")
	var pos BoardPos

	if game, ok = games[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	if err = c.BindJSON(&pos); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = game.flagCell(pos.Row, pos.Col); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}
