package router

import (
	"danjian/consts"
	"danjian/util"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Summary 团队风采
// @Produce json
// @Param page query string true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /Get [get]
func Get(c *gin.Context) {
	appG := util.Gin{C: c}

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"name": "nihao",
	})
}