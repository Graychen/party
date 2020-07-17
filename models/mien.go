package models

import (
	"danjian/consts"
	orm "danjian/database"
)

type Mien struct {
	Content    string `gorm:"column:content"`
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	Num        int    `gorm:"column:num"`
	Status     int    `gorm:"column:status"`
	Title      string `gorm:"column:title"`
	Type       int    `gorm:"column:type"`
}

var Miens []Mien

//列表
func (mien *Mien) List(page int, mienType int) (miens []Mien, err error) {
	model := orm.Eloquent
	if mienType != 0 {
		model = model.Where("type = ?", mienType)
	}
	if err = model.Offset((page - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&miens).Error; err != nil {
		return
	}
	return
}

//创建
func (mien *Mien) Create(obj interface{}) (activities []Mien, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}
