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

type theoy struct {
	Title      string `json:"Title" form:"Title" valid:"Required; MaxSize(30)"`
	Content    string `json:"Content" form:"Content" valid:"Required; MaxSize(500)"`
	Num        string `json:"Num" form:"Num"  valid:"Required; MaxSize(30)"`
	CreateTime string `json:"CreateTime" form:"CreateTime"  valid:"Required; MaxSize(30)"`
	ListImgUrl string `json:"ListImgUrl" form:"ListImgUrl" valid:"Required; MaxSize(30)"`
	Status     string `json:"Status" form:"Status" valid:"Required; MaxSize(30)"`
}

// @Summary 理论分享列表
// @Produce json
// @Param Page query int true "page"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/theoies [get]
func Theoies(c *gin.Context) {
	var theoy model.Theoy
	appG := util.Gin{C: c}
	page := c.Query("Page")
	pageNumber, _ := strconv.Atoi(page)

	result, err := theoy.List(pageNumber)

	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 理论详情
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/theoies/{id} [get]
func Theoy(c *gin.Context) {
	var theoy model.Theoy
	appG := util.Gin{C: c}
	id := c.Param("id")
	IdNumber, _ := strconv.Atoi(id)

	result, err := theoy.First(IdNumber)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Summary 添加理论分享
// @Accept multipart/form-data
// @Produce json
// @Param Title formData string true "理论标题"
// @Param Content formData string true "理论内容"
// @Param Status formData int true "1正常显示,0不显示"
// @Param ListImgUrl formData string true "封面图片url"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"请求参数错误"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"添加文章失败"}"
// @Router /api/v1/theoies [post]
func CreateTheoy(c *gin.Context) {
	var theoyModel model.Theoy
	valid := validation.Validation{}
	appG := util.Gin{C: c}
	title := c.PostForm("Title")
	content := c.PostForm("Content")
	listImgUrl := c.PostForm("ListImgUrl")
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	status := "1"
	num := "0"
	param := &theoy{Title: title, Content: content, Status: status, Num: num, CreateTime: nowString, ListImgUrl: listImgUrl}
	ok, _ := valid.Valid(param)
	if !ok {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}
	result, err := theoyModel.Create(param)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusCreated, consts.SUCCESS, result)
}
