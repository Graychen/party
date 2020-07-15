package models

import (
	orm "danjian/database"
	    "danjian/consts"
)

type Activity struct {
	ID  int64     `gorm:"column:id"`
	Theme  string `gorm:"column:theme"`
	Time  int64 `gorm:"column:time"`
	Address  string `gorm:"column:address"`
	Content  string `gorm:"column:content"`
	Totol  int64 `gorm:"column:totol"`
	Status  int64 `gorm:"column:status"`
	CreateTime  int64 `gorm:"column:create_time"`
}

var Activities []Activity

//列表
func (activity *Activity) List(page int)(activities []Activity, err error) {
	if err = orm.Eloquent.Offset((page-1)*consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&activities).Error; err != nil {
		return 
	}
	return
}