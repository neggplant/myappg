package routers

import (
	"myappg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userCtrl := controllers.NewUserController()

	api := r.Group("/api")
	{
		api.GET("/users", userCtrl.GetUsers)
		api.POST("/users", userCtrl.CreateUser)
		api.GET("/users/:id", userCtrl.GetUserByID)
		api.PUT("/users/:id", userCtrl.UpdateUser)
		api.DELETE("/users/:id", userCtrl.DeleteUser)
	}

	return r
}
