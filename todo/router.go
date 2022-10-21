package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
	"github.com/hugosilvanunes/mytodos/server/middleware"
)

func Router(r *gin.Engine, cfg config.Config) {
	tr := r.Group("/todos", middleware.Auth(cfg))
	tr.GET("", FindAll(cfg))
}
