package endpoints

import (
	"database/sql"
	"depogunabangunan/apps/handler"

	"github.com/gin-gonic/gin"
)

// InitializeEndpoints initializes all the endpoints
func InitializeEndpoints(router *gin.RouterGroup) {
	authGroup := router.Group("/api")
	{
		authGroup.POST("/login", func(c *gin.Context) {
			db := c.MustGet("db").(*sql.DB)
			handler.LoginHandler(c, db)
		})
	}
}
