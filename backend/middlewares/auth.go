package middlewares

import (
	"github.com/easydaniel/werewolves/backend/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username, ok := session.Get("user").(string)
		if ok {
			var user models.User
			user.Username = username
			db.Where(user).Find(&user)
			c.Set("user", &user)
		}
		c.Next()
	}
}

func GetUser(c *gin.Context) (*models.User, error) {
	user, ok := c.Get("user")
	if !ok {
		return nil, nil
	}
	modelUser, ok := user.(*models.User)
	if !ok {
		return nil, nil
	}
	return modelUser, nil
}
