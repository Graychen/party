package main

import (
	"fmt"
	"github.com/detectiveHLH/go-backend-starter/router"
	"github.com/detectiveHLH/go-backend-starter/setting"
	"github.com/fvbock/endless"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"syscall"
)

type User struct {
	gorm.Model
	Name string
	Password string
}

func main() {
	r := router.InitRouter()
	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
	db, err := gorm.Open("mysql", "root:206065@tcp(127.0.0.1:3306)/part?charset=utf8&parseTime=True&loc=Loca");
		db.DB().SetMaxIdleConns(10)
			db.DB().SetMaxOpenConns(100)
		defer db.Close()
}