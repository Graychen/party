package api

import (
	"danjian/consts"
	"danjian/util"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary 上传图片
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传图片"
// @Success 201 {string} json "{"code":201,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":null,"msg":"图片验证失败"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"上传失败"}"
// @Router /api/v1/images [post]
func CreateImages(c *gin.Context) {
	appG := util.Gin{C: c}
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	if err != nil {
		appG.Response(http.StatusBadRequest, consts.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	out, err := os.Create("uploads/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		appG.Response(consts.ERROR, consts.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	filepath := "http://58.87.121.2/uploads/" + filename
	appG.Response(http.StatusCreated, consts.SUCCESS, filepath)
}

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
