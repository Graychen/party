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

type apply struct {
	ActivityId string `json:"ActivityId" form:"ActivityId" valid:"Required"`
	UserId     string `json:"UserId" form:"UserId" valid:"Required"`
	Status     string `json:"Status" form:"Status" valid:"Required"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required"`
}

// @Summary 活动报名列表
// @Produce json
// @Param Page query int true "Page"
// @Param ActivityId query int false "活动id"
// @Param UserId query int false "用户id"
// @Param Status query int false "1通过,0不通过"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/applies [get]
func Applies(c *gin.Context) {
	var apply model.Apply
	appG := util.Gin{C: c}
	page := c.Query("Page")
	activityId := c.Query("ActivityId")
	userId := c.Query("UserId")
	status := c.Query("Status")
	pageNumber, _ := strconv.Atoi(page)
	activityIdNumber, _ := strconv.Atoi(activityId)
	userIdNumber, _ := strconv.Atoi(userId)

	result, err := apply.List(pageNumber, activityIdNumber, userIdNumber, status)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 申请活动报名
// @Accept multipart/form-data
// @Produce json
// @Param ActivityId formData string true "活动id"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/applies [post]
func CreateApply(c *gin.Context) {
	var applyModel model.Apply
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	activity_id := c.PostForm("ActivityId")
	user_id := "1"
	status := "1"
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	param := &apply{ActivityId: activity_id, UserId: user_id, Status: status, CreateTime: nowString}
	ok, err := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
		/*
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		*/
	}
	result, err := applyModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
