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
	Address    string `json:"Address" form:"Address" valid:"Required"`
	Content    string `json:"Content" form:"Content" valid:"Required"`
	Joined     string `json:"Joined" form:"Joined"  valid:"Required"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required"`
	ListImgUrl string `json:"ListImgUrl" form:"ListImgUrl"`
	Status     string `json:"Status" form:"Status" valid:"Required"`
	Name       string `json:"Name" form:"Name"  valid:"Required"`
	Record     string `json:"Record" form:"Record"  valid:"Required"`
	Time       string `json:"Time" form:"Time" valid:"Required"`
}

type putDay struct {
	Id         string `json:"Id" form:"Id"`
	Address    string `json:"Address" form:"Address"`
	Content    string `json:"Content" form:"Content"`
	Joined     string `json:"Joined" form:"Joined"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required"`
	ListImgUrl string `json:"ListImgUrl" form:"ListImgUrl"`
	Status     string `json:"Status" form:"Status"`
	Name       string `json:"Name" form:"Name"`
	Record     string `json:"Record" form:"Record"`
	Time       string `json:"Time" form:"Time"`
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

// @Summary 修改主题党日动态
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Id"
// @Param Address formData string false "地点"
// @Param Content formData string false "活动内容"
// @Param Name formData string false "活动名称"
// @Param ListImgUrl formData string false "封面图片url"
// @Param Joined formData string false "参加对象"
// @Param Record formData string false "记录人"
// @Param Time formData string false "时间"
// @Success 200 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 404 {string} json "{"code":404,"data":null,"msg":"没有这个资源"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/days/{id} [put]
func UpdateDay(c *gin.Context) {
	var dayModel model.Day
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	address := c.PostForm("Address")
	content := c.PostForm("Content")
	name := c.PostForm("Name")
	status := c.PostForm("Status")
	listImgUrl := c.PostForm("ListImgUrl")
	joined := c.PostForm("Joined")
	record := c.PostForm("Record")
	dayTime := c.PostForm("Time")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	id := c.Param("id")
	param := &putDay{Id: id, Time: dayTime, Joined: joined, Record: record, Address: address, Content: content, Name: name, Status: status, CreateTime: nowString, ListImgUrl: listImgUrl}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := dayModel.Update(id, dayTime, joined, record, address, content, name, status, nowString, listImgUrl)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 删除主题党日动态
// @Produce json
// @Param id path int true "Id"
// @Success 204 {string} json "{"code":204,"data":{},"msg":"ok"}"
// @Router /api/v1/days/{id} [delete]
func DeleteDay(c *gin.Context) {
	var dayModel model.Day
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)
	result, err := dayModel.Delete(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusNoContent, consts.SUCCESS, result)
}
