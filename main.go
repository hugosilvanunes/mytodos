package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
	"github.com/hugosilvanunes/mytodos/user"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	user.Router(r, *cfg)

	if err := r.Run(cfg.ENV.Port); err != nil {
		log.Fatal(err)
	}
}
