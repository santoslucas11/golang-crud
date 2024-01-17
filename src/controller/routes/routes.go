package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santoslucas11/goland-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	//get by id
	r.GET("/get-user-id/:userId", controller.FindUserById)
	//get by email
	r.GET("/get-user-email/:userEmail", controller.FindUserByEmail)
	//insert
	r.POST("/create-user", controller.CreateUser)
	//update
	r.PUT("/update-user/:userId", controller.UpdateUser)
	//delete
	r.DELETE("/delete-user/:userId", controller.DeleteUser)
}
