package models

import (
	orm "danjian/database"
	    "danjian/consts"
)

type Theoy struct {
        Content    string   `gorm:"column:content"`
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        Num        int      `gorm:"column:num"`
        Status     int      `gorm:"column:status"`
        Title      string   `gorm:"column:title"`
}

var Theoies []Theoy

//列表
func (theoy *Theoy) List(page int)(theoies []Theoy, err error) {
	if err = orm.Eloquent.Offset((page-1)*consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&theoies).Error; err != nil {
		return 
	}
	return
}