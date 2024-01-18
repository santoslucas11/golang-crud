package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/santoslucas11/goland-crud/src/configuration/validation"
	"github.com/santoslucas11/goland-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("[BEGIN] - Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("[ERROR] - Error trying to marshal object - error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
