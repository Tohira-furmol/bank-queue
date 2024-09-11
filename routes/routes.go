package routes

import (
	"github.com/gin-gonic/gin"
	"bank-queue-system/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.RegisterUser)
	r.POST("/book", controllers.BookQueue)

	return r
}


