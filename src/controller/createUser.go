package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/santoslucas11/goland-crud/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
