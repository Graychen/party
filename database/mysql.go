package database

import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
    "github.com/jinzhu/gorm"
    "fmt"
)

var Eloquent *gorm.DB

func init() {
	var err error
    Eloquent, err = gorm.Open("mysql", "root:206065@tcp(127.0.0.1:3306)/party?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	Eloquent.SingularTable(true) //禁用复数形式表名
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {//加上表前缀
		return "party_" + defaultTableName;
	}
    Eloquent.LogMode(true)//调试时打印sql语句
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Eloquent.Error != nil {
        fmt.Printf("database error %v", Eloquent.Error)
    }
}