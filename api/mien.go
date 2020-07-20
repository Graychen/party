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

type mien struct {
	Type       string `json:"Type" form:"Type" valid:"Required"`
	Title      string `json:"Title" form:"Title" valid:"Required"`
	Content    string `json:"Content" form:"Content" valid:"Required"`
	Num        string `json:"Num" form:"Num"  valid:"Required"`
	ListImgUrl string `json:"ListImgUrl" form:"ListImgUrl" valid:"Required"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required"`
	Status     string `json:"Status" form:"Status" valid:"Required"`
}

// @Summary 风采列表
// @Produce json
// @Param Type query int false "类型1团队风采,2活动风采"
// @Param Page query int true "Page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/miens [get]
func Miens(c *gin.Context) {
	var mien model.Mien
	appG := util.Gin{C: c}
	mienType := c.Query("Type")
	page := c.Query("Page")
	pageNumber, _ := strconv.Atoi(page)
	typeNumber, _ := strconv.Atoi(mienType)

	result, err := mien.List(pageNumber, typeNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 风采详情
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/miens/{id} [get]
func Mien(c *gin.Context) {
	var mien model.Mien
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := mien.First(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 添加风采
// @Accept multipart/form-data
// @Produce json
// @Param Type formData string true "类型1团队风采,2活动风采"
// @Param Title formData string true "风采标题"
// @Param Content formData string true "风采内容"
// @Param ListImgUrl formData string true "封面图片url"
// @Param Status formData int true "1正常显示,0不显示"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/miens [post]
func CreateMien(c *gin.Context) {
	var mienModel model.Mien
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	mineType := c.PostForm("Type")
	title := c.PostForm("Title")
	content := c.PostForm("Content")
	listImgUrl := c.PostForm("ListImgUrl")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	status := "1"
	num := "0"
	param := &mien{Type: mineType, Title: title, Content: content, Status: status, Num: num, CreateTime: nowString, ListImgUrl: listImgUrl}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := mienModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
