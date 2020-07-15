package router

import (
	  "danjian/consts"
	  "danjian/util"
	  "github.com/gin-gonic/gin"
	  "net/http"
model "danjian/models"
	  "strconv"
)

// @Summary 理论分享列表
// @Produce json
// @Param page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/theoies [get]
func Theoies(c *gin.Context) {
	var theoy model.Theoy
	appG := util.Gin{C: c}
	page := c.Query("page")
	pageNumber , _ := strconv.Atoi(page)

	result, err := theoy.List(pageNumber)
	
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}