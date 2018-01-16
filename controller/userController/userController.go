package userController

import (
	"github.com/gin-gonic/gin"
	"sauth/service/userService"
)

func Add(ctx *gin.Context) {
	res, err := userService.Add(ctx)
	ctx.JSON(200, gin.H{
		"error": err,
		"data":  res,
	})
}
