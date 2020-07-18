package models

import (
	"danjian/consts"
	orm "danjian/database"
	"strconv"
)

type Apply struct {
	ActivityID int `gorm:"column:activity_id"`
	CreateTime int `gorm:"column:create_time"`
	ID         int `gorm:"column:id;primary_key"`
	Status     int `gorm:"column:status"`
	UserID     int `gorm:"column:user_id"`
}

//列表
func (apply *Apply) List(pageNumber int, activityIdNumber int, userIdNumber int, status string) (appies []Apply, err error) {
	model := orm.Eloquent
	if activityIdNumber != 0 {
		model = model.Where("activity_id = ?", activityIdNumber)
	}
	if userIdNumber != 0 {
		model = model.Where("user_id = ?", userIdNumber)
	}
	if status != "" {
		statusNumber, _ := strconv.Atoi(status)
		model = model.Where("status = ?", statusNumber)
	}
	if err = model.Offset((pageNumber - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&appies).Error; err != nil {
		return
	}
	return
}

//创建
func (apply *Apply) Create(obj interface{}) (appies []Apply, err error) {
	if err = orm.Eloquent.Create(obj).Error; err != nil {
		return
	}
	return
}
