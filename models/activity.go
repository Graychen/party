package models

import (
	"danjian/consts"
	orm "danjian/database"
)

type Activity struct {
	ID         int    `gorm:"column:id;primary_key"`
	Theme      string `gorm:"column:theme"`
	Time       int    `gorm:"column:time"`
	Address    string `gorm:"column:address"`
	Content    string `gorm:"column:content"`
	Totol      int    `gorm:"column:totol"`
	Status     int    `gorm:"column:status"`
	CreateTime int    `gorm:"column:create_time"`
}

// TableName sets the insert table name for this struct type
func (t *Activity) TableName() string {
	return "party_activity"
}

var Activities []Activity

//列表
func (activity *Activity) List(page int) (activities []Activity, err error) {
	if err = orm.Eloquent.Offset((page - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&activities).Error; err != nil {
		return
	}
	return
}

//创建
func (activity *Activity) Create(obj interface{}) (activities []Activity, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}
