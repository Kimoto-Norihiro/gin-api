package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/models"
)

type UserController struct{}

// NewHealthController is constructer for HealthController
func NewUserController() *UserController {
	return new(UserController)
}

// Index is index route for health
func (uc *UserController) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": http.StatusText(http.StatusOK),
			"data": user,
	})
}
