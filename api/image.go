package api

import (
	"danjian/consts"
	"danjian/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 上传图片
// @Accept multipart/form-data
// @Produce json
// @Param File formData file true "上传图片"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"图片验证失败"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"上传失败"}"
// @Router /api/v1/images [post]
func CreateImages(c *gin.Context) {
	appG := util.Gin{C: c}
	file, err := c.FormFile("File")
	if err != nil {
		appG.Response(http.StatusBadRequest, consts.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	file.Filename = now + "_" + file.Filename
	if err := c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
		appG.Response(consts.ERROR, consts.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	filepath := "uploads/" + file.Filename
	appG.Response(http.StatusCreated, consts.SUCCESS, filepath)
}

/*
func CreateImages(c *gin.Context) {
	appG := util.Gin{C: c}
	file, header, err := c.Request.FormFile("File")
	filename := header.Filename
	if err != nil {
		appG.Response(http.StatusBadRequest, consts.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	newFileName := now + "_" + filename
	out, err := os.Create("uploads/" + newFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	filepath := "uploads/" + newFileName
	appG.Response(http.StatusCreated, consts.SUCCESS, filepath)
}
*/

func Images(c *gin.Context) {
	appG := util.Gin{C: c}
	imageName := c.Query("imageName")
	path := "../uploads/" + imageName
	info, err := ioutil.ReadFile(path)
	if err != nil {
		appG.Response(http.StatusNotFound, consts.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	fmt.Println(info)
}

// @Summary 删除图片
// @Accept multipart/form-data
// @Produce json
// @Param ImageName query string true "活动主题"
// @Success 204 {string} json "{"code":204,"data":null,"msg":"删除成功"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"删除失败"}"
// @Router /api/v1/images [delete]
func DeleteImages(c *gin.Context) {
	appG := util.Gin{C: c}
	imageName := c.Query("ImageName")
	del := os.Remove("uploads/" + imageName)
	if del != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_DELETE_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusNoContent, consts.SUCCESS, nil)
}
