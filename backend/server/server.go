package server

import (
	"github.com/gin-gonic/gin"

	"github.com/easydaniel/werewolves/backend/controllers"
)

func Run() {
	router := gin.Default()

	gameController := controllers.NewGameController()

	game := router.Group("/games")
	game.GET("/", gameController.Gets)
	game.POST("/", gameController.Post)

	router.Run(":8081")
}
