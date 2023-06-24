package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/models"
	"github.com/Kimoto-Norihiro/gin-api/usecases"
)

type UserHandler struct{
	usecase usecases.Usecase
}

func NewUserHandler(usecase usecases.Usecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	var err error
	
	if err = c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = uh.usecase.CreateUser(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "success"})
}
