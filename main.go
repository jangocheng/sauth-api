package main

import (
	"github.com/gin-gonic/gin"
	"sauth/router"
	"sauth/util"
	"sauth/controller/userController"
)

func main() {
	gin.SetMode(gin.ReleaseMode)              // 发布模式
	engine := gin.Default()                   // web engine
	engine.Use(util.Cors())                   // 跨域
	engine.Use(userController.Authenticate()) // 权限认证
	router.Router(engine)                     // 路由
	engine.Run()
}
