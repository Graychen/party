package main

import (
	_ "danjian/database"
	orm "danjian/database"
	"danjian/router"
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func main() {
	r := router.InitRouter()
	address := fmt.Sprintf("%s:%s", "0.0.0.0", "80")
	log.Printf("address is %s", address)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	defer orm.Eloquent.Close()
}
