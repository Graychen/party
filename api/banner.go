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

type banner struct {
	ImgUrl     string `json:"ImgUrl" form:"ImgUrl" valid:"Required"`
	Type       string `json:"Type" form:"Type" valid:"Required"`
	Status     string `json:"Status" form:"Status" valid:"Required"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required"`
	UserId     string `json:"UserId" form:"UserId"  valid:"Required"`
}

// @Summary 首页封面列表
// @Produce json
// @Param page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/banners [get]
func Banners(c *gin.Context) {
	var banner model.Banner
	appG := util.Gin{C: c}
	page := c.Query("Page")
	pageNumber, _ := strconv.Atoi(page)

	result, err := banner.List(pageNumber)

	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 添加封面图片
// @Accept multipart/form-data
// @Produce json
// @Param ImgUrl formData string true "图片url"
// @Param Type formData int true "1：党建工作室 2：活动报名 3：理论分享 4：主题党日 5：党建动态"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/banners [post]
func CreateBanner(c *gin.Context) {
	var bannerModel model.Banner
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	imgUrl := c.PostForm("ImgUrl")
	bannerType := c.PostForm("Type")
	status := "1"
	userId := "1"
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	param := &banner{ImgUrl: imgUrl, Type: bannerType, Status: status, CreateTime: nowString, UserId: userId}
	/*
		if bannerModel.Count() > 5,err==nil {
			appG.Response(http.StatusBadRequest, consts.ERROR_BANNER_FIVE, nil)
			return
		}
	*/
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := bannerModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}

// @Summary 删除封面
// @Produce json
// @Param id path int true "Id"
// @Success 204 {string} json "{"code":204,"data":{},"msg":"ok"}"
// @Router /api/v1/banners/{id} [delete]
func DeleteBanner(c *gin.Context) {
	var bannerModel model.Banner
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := bannerModel.Delete(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusNoContent, consts.SUCCESS, result)
}
