package routes

import (
	"service-golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, userHandler *handlers.UserHandler){
	api := r.Group("/api/v1")
	{
		api.POST("/user", userHandler.CreateUser)
		api.GET("/users", userHandler.GetAllUser)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}
}