package router

import (
	"github.com/gin-gonic/gin"
	"sauth/controller"
	"sauth/controller/appController"
	"sauth/controller/userController"
	"sauth/controller/roleController"
	"sauth/controller/functionController"
	"sauth/controller/roleFunctionController"
	"sauth/controller/orgController"
)

/*
	访问路由
*/
func Router(engine *gin.Engine) {

	v1 := engine.Group("auth/api/v1") // 路由分组
	{
		// ping
		v1.GET("/ping", controller.Ping)

		// 应用
		v1.GET("/app/all", appController.All)
		v1.GET("/app/find", appController.Find)
		v1.POST("app/save", appController.Save)
		v1.POST("/app/delete", appController.Delete)
		v1.GET("/app/options", appController.Options)

		// 用户
		v1.POST("/user/add", userController.Add)
		v1.POST("/user/login", userController.Login)
		v1.POST("/user/find", userController.Find)

		// 角色
		v1.GET("/role/all", roleController.All)
		v1.GET("/role/find", roleController.Find)
		v1.POST("/role/save", roleController.Save)
		v1.POST("/role/delete", roleController.Delete)

		// 功能
		v1.GET("/function/all", functionController.All)
		v1.GET("/function/find", functionController.Find)

		// 角色-功能
		v1.POST("/roleFunction/save", roleFunctionController.Save)
		v1.GET("/roleFunction/find", roleFunctionController.Find)

		// 机构
		v1.GET("/org/find", orgController.Find)
		v1.POST("/org/save", orgController.Save)
		v1.POST("/org/delete", orgController.Delete)
	}
}
