package Routes

import (
	"github.com/aditya/mvcAPI/controllers"
	"github.com/gin-gonic/gin"
)

// Setup Router for APIs

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserByID)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}
