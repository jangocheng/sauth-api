package controller

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    "S_AuthVerify_01",
		"message": "用户权限认证通过",
		"data":    nil,
		"err":     nil,
	})
}
