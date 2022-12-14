package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
)

func Router(r *gin.Engine, cfg config.Config) {
	r.POST("/register", Register(cfg))
	r.POST("/login", Login(cfg))
}
