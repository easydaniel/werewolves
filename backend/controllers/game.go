package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/easydaniel/werewolves/backend/game"
	"github.com/easydaniel/werewolves/backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameController struct {
	db    *gorm.DB
	games map[string]*game.Game
}

func NewGameController(db *gorm.DB) *GameController {
	ctrl := new(GameController)
	ctrl.db = db
	ctrl.games = make(map[string]*game.Game)
	return ctrl
}

func (ctrl *GameController) GetBoardType(c *gin.Context) {
	files, err := ioutil.ReadDir("./data/boards")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Read Files Error",
		})
		return
	}

	result := []string{}

	for _, f := range files {
		basename := f.Name()
		result = append(result, strings.TrimSuffix(basename, filepath.Ext(basename)))
	}

	c.JSON(http.StatusOK, result)
}

// func (ctrl *GameController) Status(c *gin.Context) {
// 	id := c.Param("gameID")
// 	c.JSON(http.StatusOK, ctrl.games[id])
// }

type GameCreateBody struct {
	Board string `json:"board"`
}

func (ctrl *GameController) Create(c *gin.Context) {
	var body GameCreateBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
	}
	user, _ := middlewares.GetUser(c)
	newGame, _ := game.NewGame(body.Board, &game.Member{Name: user.Name})
	ctrl.games[newGame.ID] = newGame

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Create Game Error",
		})
	}

	c.JSON(http.StatusOK, newGame)
}

type GameJoinBody struct {
	GameID string `json:"game_id"`
}

func (ctrl *GameController) JoinRoom(c *gin.Context) {
	var body GameJoinBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[body.GameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[body.GameID]
	user, _ := middlewares.GetUser(c)
	err = g.JoinRoom(&game.Member{Name: user.Name})

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type GameExitRoomBody struct {
	GameID string `json:"game_id"`
}

func (ctrl *GameController) ExitRoom(c *gin.Context) {
	var body GameExitRoomBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[body.GameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[body.GameID]
	user, _ := middlewares.GetUser(c)
	err = g.ExitRoom(&game.Member{Name: user.Name})

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type GameSetSeatBody struct {
	GameID string `json:"game_id"`
	Seat   int    `json:"seat"`
}

func (ctrl *GameController) SetSeat(c *gin.Context) {
	var body GameSetSeatBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[body.GameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[body.GameID]
	user, _ := middlewares.GetUser(c)
	err = g.SetSeat(&game.Member{Name: user.Name}, body.Seat)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type GameExitSeatBody struct {
	GameID string `json:"game_id"`
	Seat   int    `json:"seat"`
}

func (ctrl *GameController) ExitSeat(c *gin.Context) {
	var body GameExitSeatBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[body.GameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[body.GameID]
	// user, _ := middlewares.GetUser(c)
	err = g.ExitSeat(body.Seat)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
