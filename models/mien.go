package models

import (
	"danjian/consts"
	orm "danjian/database"
	"strconv"
)

type Mien struct {
	Content    string `gorm:"column:content"`
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	Num        int    `gorm:"column:num"`
	Status     int    `gorm:"column:status"`
	ListImgUrl string `gorm:"column:list_img_url"`
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
func (mien *Mien) Create(obj interface{}) (miens []Mien, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}

//详情
func (mien *Mien) First(id int) (miens []Mien, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&miens).Error; err != nil {
		return
	}
	return
}

//修改
func (mien *Mien) Update(id string, mienType string, title string, content string, status string, nowString string, listImgUrl string) (miens []Mien, err error) {
	IdNumber, _ := strconv.Atoi(id)
	nowNumber, _ := strconv.Atoi(nowString)
	statusNumber, _ := strconv.Atoi(status)
	typeNumber, _ := strconv.Atoi(mienType)
	model := orm.Eloquent
	if err = model.First(&mien, IdNumber).Error; err != nil {
		return
	}
	if mienType != "" {
		mien.Type = typeNumber
	}
	if title != "" {
		mien.Title = title
	}
	if content != "" {
		mien.Content = content
	}
	if nowString != "" {
		mien.CreateTime = nowNumber
	}
	if status != "" {
		mien.Status = statusNumber
	}
	if listImgUrl != "" {
		mien.ListImgUrl = listImgUrl
	}
	if err = model.Save(&mien).Error; err != nil {
		return
	}
	return
}

//删除
func (mien *Mien) Delete(id int) (miens []Mien, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.Delete(&mien).Error; err != nil {
		return
	}
	return
}
