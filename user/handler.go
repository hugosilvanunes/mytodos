package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
	"golang.org/x/crypto/bcrypt"
)

func Register(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var register RegisterInput
		if err := c.ShouldBindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }

        register.Password = string(hashedPassword)

        _, err = cfg.DB.NamedExec(FIND_USER_BY_ID, register)
        if err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	       return
        }

        c.Status(http.StatusCreated)
	}
}

func Login(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login LoginInput
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
