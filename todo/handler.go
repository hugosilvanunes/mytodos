package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/mytodos/config"
)

func FindAll(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.MustGet("userID").(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userID must be a string"})
			return
		}

		users := make([]Todo, 0)
		if err := cfg.DB.Select(&users, FIND_TODOS_BY_USER_ID_SQL, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
