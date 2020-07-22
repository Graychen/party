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
	ListImgUrl string `gorm:"column:list_img_url"`
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

//修改
func (dynamic *Dynamic) Update(id string, title string, content string, status string, num string, nowString string, listImgUrl string) (dynamics []Dynamic, err error) {
	IdNumber, _ := strconv.Atoi(id)
	numNumber, _ := strconv.Atoi(num)
	dynamicTimeNumber, _ := strconv.Atoi(nowString)
	statusNumber, _ := strconv.Atoi(status)
	model := orm.Eloquent
	if err = model.First(&dynamic, IdNumber).Error; err != nil {
		return
	}
	if title != "" {
		dynamic.Title = title
	}
	if content != "" {
		dynamic.Content = content
	}
	if num != "" {
		dynamic.Num = numNumber
	}
	if content != "" {
		dynamic.Content = content
	}
	if status != "" {
		dynamic.Status = statusNumber
	}
	if nowString != "" {
		dynamic.CreateTime = dynamicTimeNumber
	}
	if listImgUrl != "" {
		dynamic.ListImgUrl = listImgUrl
	}
	if err = model.Save(&dynamic).Error; err != nil {
		return
	}
	return
}

//删除
func (dynamic *Dynamic) Delete(id int) (dynamics []Dynamic, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.Delete(&dynamic).Error; err != nil {
		return
	}
	return
}
