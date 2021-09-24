package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xtareq/go_api/config"
	"github.com/xtareq/go_api/controllers"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.DbConnection()
	authController controllers.AuthController = controllers.NewAuthController()
)

func main() {
	defer config.CloseDb(db)
	r := gin.Default()
	/* 	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	}) */

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
