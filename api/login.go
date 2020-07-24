package api

import (
	"danjian/consts"
	model "danjian/models"
	"danjian/service/authentication"
	"danjian/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 登录
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login [get]
func Login(c *gin.Context) {
	appG := util.Gin{C: c}
	valid := validation.Validation{}
	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		appG.Response(http.StatusOK, consts.INVALID_PARAMS, nil)
		return
	}
	authService := authentication.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusOK, consts.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary 获取accessToken
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /userId [get]
func UserInfo(c *gin.Context) {
	var user model.Users
	appG := util.Gin{C: c}
	token, err := user.GetToken()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_TOKEN, nil)
		return
	}
	code := c.Query("code")
	userInfo, err := user.GetUserId(token, code)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, string(userInfo))
}
