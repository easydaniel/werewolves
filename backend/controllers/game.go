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

func (ctrl *GameController) Status(c *gin.Context) {
	gameID := c.Param("gameID")
	user, _ := middlewares.GetUser(c)
	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}
	c.JSON(http.StatusOK, ctrl.games[gameID].Status(&game.Member{Name: user.Name}))
}

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

func (ctrl *GameController) JoinRoom(c *gin.Context) {
	gameID := c.Param("gameID")

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[gameID]
	user, _ := middlewares.GetUser(c)
	err := g.JoinRoom(&game.Member{Name: user.Name})

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

func (ctrl *GameController) ExitRoom(c *gin.Context) {
	gameID := c.Param("gameID")

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[gameID]
	user, _ := middlewares.GetUser(c)
	err := g.ExitRoom(&game.Member{Name: user.Name})

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
	Seat int `json:"seat"`
}

func (ctrl *GameController) SetSeat(c *gin.Context) {
	gameID := c.Param("gameID")

	var body GameSetSeatBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[gameID]
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
	Seat int `json:"seat"`
}

func (ctrl *GameController) ExitSeat(c *gin.Context) {
	gameID := c.Param("gameID")

	var body GameExitSeatBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[gameID]
	user, _ := middlewares.GetUser(c)
	seat := -1
	if !g.IsHost(&game.Member{Name: user.Name}) {
		seat, err = g.GetSeat(&game.Member{Name: user.Name})
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
	}
	err = g.ExitSeat(seat)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

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

func (ctrl *GameController) TestStart(c *gin.Context) {
	gameID := c.Param("gameID")

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	ctrl.games[gameID].FillTestUser()

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type GameKillBody struct {
	Seat int `json:"seat"`
}

func (ctrl *GameController) Kill(c *gin.Context) {
	gameID := c.Param("gameID")
	var body GameKillBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	if _, ok := ctrl.games[gameID]; !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No Game",
		})
		return
	}

	g := ctrl.games[gameID]
	user, _ := middlewares.GetUser(c)
	if !g.IsHost(&game.Member{Name: user.Name}) {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"error": "Permission denied",
		})
		return
	}

	g.Kill(body.Seat)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
