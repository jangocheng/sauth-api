package userService

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"sauth/model/userModel"
	"sauth/util"
)

func Add(ctx *gin.Context) (sql.Result, error) {
	var model userModel.User
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}
	model.Password = util.Md5Encrypt(model.Password) // 加密
	rows, err := userModel.Insert(&model)
	return rows, err
}
