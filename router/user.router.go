package router

import (
	"github.com/Bassiouni/money-mate-api/utils"
)

func BindUserRouter() {
	var userRouter = RootRouter.Group("/users")

	userRouter.POST("/new", utils.CreateUser)
	userRouter.GET("/", utils.GetAllUsers)
	userRouter.GET("/:id", utils.GetUserByID)
	// userRouter.GET("/:email", utils.GetUserByEmail)
}
