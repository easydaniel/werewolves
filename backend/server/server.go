package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/easydaniel/werewolves/backend/logp"
	"github.com/easydaniel/werewolves/backend/ws"
)

func Run() {
	logp.Init()

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	ws := ws.NewWebSocket()
	// router.Use(ginzap.Ginzap(logp.L(), time.RFC3339, true))
	// router.Use(ginzap.RecoveryWithZap(logp.L(), true))

	router.GET("/ws", gin.WrapF(ws.Handler))

	logp.L().Info("Listen", zap.Int("port", 8081))
	router.Run("0.0.0.0:8081")
}
