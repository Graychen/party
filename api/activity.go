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

type activity struct {
	Theme      string `json:"Theme" form:"Theme" valid:"Required; MaxSize(100)"`
	Time       string `json:"Time" form:"Time" valid:"Required; MaxSize(30)"`
	Address    string `json:"Address" form:"Address" valid:"Required; MaxSize(30)"`
	Content    string `json:"Content" form:"Content" valid:"Required; MaxSize(500)"`
	Totol      string `json:"Totol" form:"Totol"  valid:"Required; MaxSize(30)"`
	Status     string `json:"Status" form:"Status" valid:"Required; MaxSize(30)"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required; MaxSize(30)"`
}

// @Summary 活动列表
// @Produce json
// @Param Page query int true "Page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/activities [get]
func Activities(c *gin.Context) {
	var activity model.Activity
	appG := util.Gin{C: c}
	page := c.Query("Page")
	pageNumber, _ := strconv.Atoi(page)

	result, err := activity.List(pageNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 活动详情
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/activities/{id} [get]
func Activity(c *gin.Context) {
	var activity model.Activity
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := activity.First(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 添加活动
// @Accept multipart/form-data
// @Produce json
// @Param Theme formData string true "活动主题"
// @Param Time formData int true "活动时间"
// @Param Address formData string true "活动地址"
// @Param Content formData string true "活动内容"
// @Param Totol formData int true "活动人数"
// @Param Status formData int true "1正常显示,0不显示"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/activities [post]
func CreateActivity(c *gin.Context) {
	var activityModel model.Activity
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	theme := c.PostForm("Theme")
	activityTime := c.PostForm("Time")
	address := c.PostForm("Address")
	content := c.PostForm("Content")
	totol := c.PostForm("Totol")
	status := c.PostForm("Status")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	param := &activity{Theme: theme, Time: activityTime, Address: address, Content: content, Totol: totol, Status: status, CreateTime: nowString}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := activityModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
