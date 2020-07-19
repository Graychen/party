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
	ListImgUrl string `gorm:"column:list_img_url"`
	Totol      int    `gorm:"column:totol"`
	Status     int    `gorm:"column:status"`
	CreateTime int    `gorm:"column:create_time"`
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

//详情
func (activity *Activity) First(id int) (activities []Activity, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&activities).Error; err != nil {
		return
	}
	return
}
