package models

import (
	"danjian/consts"
	orm "danjian/database"
	"strconv"
)

type Dynamic struct {
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	Status     int    `gorm:"column:status"`
	Num        int    `gorm:"column:num"`
}

var Dynamics []Dynamic

//列表
func (dynamic *Dynamic) List(page int, status string) (dynamics []Dynamic, err error) {
	model := orm.Eloquent
	if status != "" {
		statusNumber, _ := strconv.Atoi(status)
		model = model.Where("status = ?", statusNumber)
	}
	if err = model.Offset((page - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&dynamics).Error; err != nil {
		return
	}
	return
}

//创建
func (dynamic *Dynamic) Create(obj interface{}) (dynamics []Dynamic, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}

//详情
func (dynamic *Dynamic) First(id int) (dynamics []Dynamic, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&dynamics).Error; err != nil {
		return
	}
	return
}
