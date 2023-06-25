package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/handlers"
	"github.com/Kimoto-Norihiro/gin-api/infra/db"
	"github.com/Kimoto-Norihiro/gin-api/repositories"
	"github.com/Kimoto-Norihiro/gin-api/usecases"
)


func InitUser(r *gin.Engine) {
	repo := repositories.NewUserRepository(db.Db)
	usecase := usecases.NewUserUsecase(repo)
	handler := handlers.NewUserHandler(usecase)

	user := r.Group("/user")
	{
		user.POST("/", handler.CreateUser)
	}
}