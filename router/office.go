package router

import (
	  "danjian/consts"
	  "danjian/util"
	  "github.com/gin-gonic/gin"
	  "net/http"
model "danjian/models"
)

// @Summary 团队风采
// @Produce json
// @Param page query string true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/offices [get]
func Get(c *gin.Context) {
	var office model.Office
	appG := util.Gin{C: c}

	result, err := office.List()

	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}