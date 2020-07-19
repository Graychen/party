package api

import (
	"danjian/consts"
	model "danjian/models"
	"danjian/util"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type dynamic struct {
	Title      string `json:"Title" form:"Title" valid:"Required; MaxSize(30)"`
	Content    string `json:"Content" form:"Content" valid:"Required; MaxSize(500)"`
	Num        string `json:"Num" form:"Num"  valid:"Required; MaxSize(30)"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required; MaxSize(30)"`
	Status     string `json:"Status" form:"Status" valid:"Required; MaxSize(30)"`
}

// @Summary 党建动态列表
// @Produce json
// @Param Page query int true "Page"
// @Param Status query int false "Status"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/dynamics [get]
func Dynamics(c *gin.Context) {
	var dynamicModel model.Dynamic
	appG := util.Gin{C: c}
	page := c.Query("Page")
	status := c.Query("Status")
	pageNumber, _ := strconv.Atoi(page)

	result, err := dynamicModel.List(pageNumber, status)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 党建动态详情
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/dynamics/{id} [get]
func Dynamic(c *gin.Context) {
	var dynamicModel model.Dynamic
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := dynamicModel.First(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 添加党建动态
// @Accept multipart/form-data
// @Produce json
// @Param Title formData string true "党建动态标题"
// @Param Content formData string true "党建动态内容"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/dynamics [post]
func CreateDynamic(c *gin.Context) {
	var dynamicModel model.Dynamic
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	title := c.PostForm("Title")
	content := c.PostForm("Content")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	status := "1"
	num := "0"
	param := &dynamic{Title: title, Content: content, Status: status, Num: num, CreateTime: nowString}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := dynamicModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
