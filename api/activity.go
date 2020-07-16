package api

import (
	"danjian/consts"
	model "danjian/models"
	"danjian/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 活动列表
// @Produce json
// @Param page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/activities [get]
func Activities(c *gin.Context) {
	var activity model.Activity
	appG := util.Gin{C: c}
	page := c.Query("page")
	pageNumber, _ := strconv.Atoi(page)

	result, err := activity.List(pageNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}
