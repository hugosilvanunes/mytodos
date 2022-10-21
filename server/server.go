package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
	"github.com/hugosilvanunes/mytodos/todo"
	"github.com/hugosilvanunes/mytodos/user"
)

func Run(cfg config.Config) error {
	r := gin.Default()

	user.Router(r, cfg)
	todo.Router(r, cfg)

	if err := r.Run(cfg.ENV.Port); err != nil {
		return err
	}

	return nil
}
