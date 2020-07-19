package models

import (
	"danjian/consts"
	orm "danjian/database"
	"strconv"
)

type Day struct {
	Address    string `gorm:"column:address"`
	Content    string `gorm:"column:content"`
	CreateTime int    `gorm:"column:create_time"`
	ID         int    `gorm:"column:id;primary_key"`
	Joined     string `gorm:"column:joined"`
	Name       string `gorm:"column:name"`
	ListImgUrl string `gorm:"column:list_img_url"`
	Record     string `gorm:"column:record"`
	Status     int    `gorm:"column:status"`
	Time       int    `gorm:"column:time"`
}

var Days []Day

//列表
func (day *Day) List(page int, status string, limit string) (days []Day, err error) {
	model := orm.Eloquent
	if status != "" {
		statusNumber, _ := strconv.Atoi(status)
		model = model.Where("status = ?", statusNumber)
	}
	pageSize := consts.PAGE_SIZE
	if limit != "" {
		pageSize, _ = strconv.Atoi(limit)
	}
	if err = model.Offset((page - 1) * pageSize).Limit(pageSize).Find(&days).Error; err != nil {
		return
	}
	return
}

//创建
func (day *Day) Create(obj interface{}) (days []Day, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}

//详情
func (day *Day) First(id int) (days []Day, err error) {
	model := orm.Eloquent
	if id != 0 {
		model = model.Where("id = ?", id)
	}
	if err = model.First(&days).Error; err != nil {
		return
	}
	return
}
