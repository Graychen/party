package main

import (
	"github.com/detectiveHLH/go-backend-starter/router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}