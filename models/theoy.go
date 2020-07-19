package models

import (
	"danjian/consts"
	orm "danjian/database"
)

type Theoy struct {
	Content    string `gorm:"column:content"`
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	Num        int    `gorm:"column:num"`
	ListImgUrl string `gorm:"column:list_img_url"`
	Status     int    `gorm:"column:status"`
	Title      string `gorm:"column:title"`
}

var Theoies []Theoy

//列表
func (theoy *Theoy) List(page int) (theoies []Theoy, err error) {
	if err = orm.Eloquent.Offset((page - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&theoies).Error; err != nil {
		return
	}
	return
}

//创建
func (theoy *Theoy) Create(obj interface{}) (theoies []Theoy, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}

//详情
func (theoy *Theoy) First(id int) (theoies []Theoy, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&theoies).Error; err != nil {
		return
	}
	return
}
