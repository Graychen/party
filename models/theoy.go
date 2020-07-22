package models

import (
	"danjian/consts"
	orm "danjian/database"
	"strconv"
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

//修改
func (theoy *Theoy) Update(id string, title string, content string, status string, nowString string, listImgUrl string) (theoies []Theoy, err error) {
	IdNumber, _ := strconv.Atoi(id)
	nowNumber, _ := strconv.Atoi(nowString)
	statusNumber, _ := strconv.Atoi(status)
	model := orm.Eloquent
	if err = model.First(&theoy, IdNumber).Error; err != nil {
		return
	}
	if title != "" {
		theoy.Title = title
	}
	if content != "" {
		theoy.Content = content
	}
	if nowString != "" {
		theoy.CreateTime = nowNumber
	}
	if status != "" {
		theoy.Status = statusNumber
	}
	if listImgUrl != "" {
		theoy.ListImgUrl = listImgUrl
	}
	if err = model.Save(&theoy).Error; err != nil {
		return
	}
	return
}

//删除
func (theoy *Theoy) Delete(id int) (theoies []Theoy, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.Delete(&theoy).Error; err != nil {
		return
	}
	return
}
