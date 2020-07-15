package router

import (
	  "danjian/consts"
	  "danjian/util"
	  "github.com/gin-gonic/gin"
	  "net/http"
model "danjian/models"
	  "strconv"
)

// @Summary 团队风采列表
// @Produce json
// @Param page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/offices [get]
func Offices(c *gin.Context) {
	var office model.Office
	appG := util.Gin{C: c}
	page := c.Query("page")
	pageNumber , _ := strconv.Atoi(page)

	result, err := office.List(pageNumber)
	
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}