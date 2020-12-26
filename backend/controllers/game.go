package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	ID    string
	Owner string
	Board string
	User  map[int]string
}

func NewGame() *Game {
	game := new(Game)
	return game
}

type GameController struct {
	games map[string]*Game
}

func NewGameController() *GameController {
	ctrl := new(GameController)
	ctrl.games = make(map[string]*Game)
	return ctrl
}

func (ctrl *GameController) Gets(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.games)
}

func (ctrl *GameController) Post(c *gin.Context) {
	type Body struct {
		Name string
	}
	var body Body
	c.BindJSON(&body)
	ctrl.games[body.Name] = new(Game)
}
