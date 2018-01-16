package userService

import (
	"github.com/gin-gonic/gin"
	"sauth/util"
	"sauth/model/userModel"
	"net/http"
)

const (
	AuthUserCookieId = "AuthUserCookieId" // sso login token's name
)

/*
	登录验证
	@return count
	"0": 验证失败  |  "1": 验证成功
*/
func LoginVerify(ctx *gin.Context) (string, error) {
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	password = util.Md5Encrypt(password)
	res, err := userModel.FindCountByUserNameAndPassword(userName, password)
	if nil != err {
		return "", err
	}
	count := string(res[0]["count"])
	if "1" == count { // 设置 cookie
		cookie := http.Cookie{
			Name:     AuthUserCookieId,
			Value:    userName,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			MaxAge:   util.SecondsInWeek,
		}
		http.SetCookie(ctx.Writer, &cookie)
	}
	return count, nil
}

/*
	权限验证
*/
func AuthVerify(ctx *gin.Context) (string, error) {
	token, err := ctx.Cookie(AuthUserCookieId)
	res, _ := userModel.FindCountByUserName(token) // 数据库验证，后面考虑用缓存验证
	count := string(res[0]["count"])
	return count, err
}
