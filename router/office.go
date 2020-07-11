package router

import (
	"github.com/astaxie/beego/validation"
	"github.com/detectiveHLH/go-backend-starter/consts"
	"github.com/detectiveHLH/go-backend-starter/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type office struct {
	Page string `valid:"Required; MaxSize(50)"`
}

// @Summary office
// @Produce json
// @Param  query string true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /offices [get]
func Get(c *gin.Context) {
	appG := util.Gin{C: c}
	valid := validation.Validation{}
	page := c.Query("page")

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"number": 1,
	})
}