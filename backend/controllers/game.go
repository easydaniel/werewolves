package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/easydaniel/werewolves/backend/game"
)

type GameController struct {
	games map[string]*game.Game
}

func NewGameController() *GameController {
	ctrl := new(GameController)
	ctrl.games = make(map[string]*game.Game)
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
	ctrl.games[body.Name] = new(game.Game)
}
