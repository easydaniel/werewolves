package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/easydaniel/werewolves/backend/middlewares"
	"github.com/easydaniel/werewolves/backend/models"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	ctrl := new(UserController)
	ctrl.db = db
	return ctrl
}

type UserRegisterBody struct {
	Name     string
	Password string
}

func (ctrl *UserController) Register(c *gin.Context) {
	var body UserRegisterBody
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	var user models.User

	// find user
	ctrl.db.Where(models.User{
		Name: body.Name,
	}).Find(&user)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "User Registered.",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	ctrl.db.Create(&models.User{
		Name:     body.Name,
		Password: string(hashedPassword),
	})

}

type UserLoginBody struct {
	Name     string
	Password string
}

func (ctrl *UserController) Login(c *gin.Context) {
	var body UserRegisterBody
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Bad Body",
		})
		return
	}

	var user models.User

	// find user
	ctrl.db.Where(models.User{
		Name: body.Name,
	}).Find(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Name or Password is wrong.",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Name or Password is wrong.",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user", body.Name)
	session.Save()
}
func (ctrl *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func (ctrl *UserController) Profile(c *gin.Context) {
	user, _ := middlewares.GetUser(c)
	c.JSON(http.StatusOK, user)
}
