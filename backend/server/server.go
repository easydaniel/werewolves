package server

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/easydaniel/werewolves/backend/config"
	"github.com/easydaniel/werewolves/backend/controllers"
	"github.com/easydaniel/werewolves/backend/logp"
	"github.com/easydaniel/werewolves/backend/middlewares"
	"github.com/easydaniel/werewolves/backend/models"
)

func Run() {
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	logp.Init()

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.User{})

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = cfg.Server.CORS.Hosts
	config.AllowCredentials = true

	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(middlewares.Auth(db))

	// router.Use(ginzap.Ginzap(logp.L(), time.RFC3339, true))
	// router.Use(ginzap.RecoveryWithZap(logp.L(), true))

	GameController := controllers.NewGameController(db)
	UserController := controllers.NewUserController(db)

	router.GET("/boardtype/", GameController.GetBoardType)
	router.POST("/games/", GameController.Create)
	// router.GET("/games/:gameID", GameController.Status)
	router.POST("/games/:gameID", GameController.JoinRoom)
	router.DELETE("/games/:gameID", GameController.ExitRoom)
	router.POST("/games/:gameID/seat/:setID", GameController.SetSeat)
	router.DELETE("/games/:gameID/seat/:setID", GameController.ExitSeat)

	router.POST("/users/register", UserController.Register)
	router.POST("/users/login", UserController.Login)
	router.POST("/users/logout", UserController.Logout)
	router.POST("/users/profile", UserController.Profile)

	logp.L().Info("Listen", zap.Int("port", cfg.Server.Port))
	router.Run(fmt.Sprintf("0.0.0.0:%v", cfg.Server.Port))
}
