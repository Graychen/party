package util

import (
	"github.com/detectiveHLH/go-backend-starter/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct{
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}){
	g.C.Json(httpCode, gin.H{
		"code": httpCode,
		"msg": consts.GetMsg(errCode),
		"data": data,
	})

	return
}