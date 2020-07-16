package api

import (
	"danjian/consts"
	model "danjian/models"
	"danjian/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 风采列表
// @Produce json
// @Param type query int false "类型1团队风采,2活动风采"
// @Param page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/miens [get]
func Miens(c *gin.Context) {
	var mien model.Mien
	appG := util.Gin{C: c}
	mienType := c.Query("type")
	page := c.Query("page")
	pageNumber, _ := strconv.Atoi(page)
	typeNumber, _ := strconv.Atoi(mienType)

	result, err := mien.List(pageNumber, typeNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}
