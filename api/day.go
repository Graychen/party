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

type day struct {
	Address    string `json:"Address" form:"Address" valid:"Required; MaxSize(30)"`
	Content    string `json:"Content" form:"Content" valid:"Required; MaxSize(500)"`
	Joined     string `json:"Joined" form:"Joined"  valid:"Required; MaxSize(30)"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required; MaxSize(30)"`
	ListImgUrl string `json:"ListImgUrl" form:"ListImgUrl" valid:"Required; MaxSize(30)"`
	Status     string `json:"Status" form:"Status" valid:"Required; MaxSize(30)"`
	Name       string `json:"Name" form:"Name"  valid:"Required; MaxSize(30)"`
	Record     string `json:"Record" form:"Record"  valid:"Required; MaxSize(30)"`
	Time       string `json:"Time" form:"Time" valid:"Required; MaxSize(30)"`
}

// @Summary 主题党日列表
// @Produce json
// @Param Page query int true "Page"
// @Param Status query int false "Status"
// @Param Limit query int false "Limit"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/days [get]
func Days(c *gin.Context) {
	var dayModel model.Day
	appG := util.Gin{C: c}
	page := c.Query("Page")
	status := c.Query("Status")
	Limit := c.Query("Limit")
	pageNumber, _ := strconv.Atoi(page)

	result, err := dayModel.List(pageNumber, status, Limit)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 主题党日详情
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/days/{id} [get]
func Day(c *gin.Context) {
	var dayModel model.Day
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := dayModel.First(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 主题党日动态
// @Accept multipart/form-data
// @Produce json
// @Param Address formData string true "地点"
// @Param Content formData string true "活动内容"
// @Param Name formData string true "活动名称"
// @Param ListImgUrl formData string true "封面图片url"
// @Param Joined formData string true "参加对象"
// @Param Record formData string true "记录人"
// @Param Time formData string true "时间"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/days [post]
func CreateDay(c *gin.Context) {
	var dayModel model.Day
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	address := c.PostForm("Address")
	content := c.PostForm("Content")
	name := c.PostForm("Name")
	listImgUrl := c.PostForm("ListImgUrl")
	joined := c.PostForm("Joined")
	record := c.PostForm("Record")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	time := c.PostForm("Time")
	status := "1"
	param := &day{Address: address, Content: content, Status: status, Joined: joined, CreateTime: nowString, ListImgUrl: listImgUrl, Record: record, Time: time, Name: name}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := dayModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
