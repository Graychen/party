package models

import (
	"danjian/consts"
	orm "danjian/database"
)

type Banner struct {
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	ImgURL     string `gorm:"column:img_url"`
	Status     int    `gorm:"column:status"`
	Type       int    `gorm:"column:type"`
	UserID     int    `gorm:"column:user_id"`
}

//列表
func (banner *Banner) List(page int) (banners []Banner, err error) {
	if err = orm.Eloquent.Offset((page - 1) * consts.BANNER_SIZE).Limit(consts.BANNER_SIZE).Find(&banners).Error; err != nil {
		return
	}
	return
}

//详情
func (banner *Banner) First(id int) (banners []Banner, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&banners).Error; err != nil {
		return
	}
	return
}

//统计
func (banner *Banner) Count() (count int, err error) {
	if err = orm.Eloquent.Count(&count).Error; err != nil {
		return
	}
	return
}

//创建
func (banner *Banner) Create(obj interface{}) (banners []Banner, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}

//删除
func (banner *Banner) Delete(id int) (banners []Banner, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.Delete(&banners).Error; err != nil {
		return
	}
	return
}
