package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/controllers"
)

func NewRouter() (r *gin.Engine, err error) {
	router := gin.Default()

  version := router.Group("/v1")

	userController := controllers.NewUserController()
	version.POST("/users", userController.Create)

	return router, nil
}