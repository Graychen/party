package models

import (
	"danjian/consts"
	orm "danjian/database"
)

type Apply struct {
	ActivityID int `gorm:"column:activity_id"`
	CreateTime int `gorm:"column:create_time"`
	ID         int `gorm:"column:id;primary_key"`
	Status     int `gorm:"column:status"`
	UserID     int `gorm:"column:user_id"`
}

//列表
func (apply *Apply) List(pageNumber int, activityIdNumber int, userIdNumber int, statusNumber int) (appies []Apply, err error) {
	model := orm.Eloquent
	if activityIdNumber != 0 {
		model = model.Where("activity = ?", activityIdNumber)
	}
	if userIdNumber != 0 {
		model = model.Where("user_id = ?", userIdNumber)
	}
	if statusNumber != 0 {
		model = model.Where("status = ?", statusNumber)
	}
	if err = model.Offset((pageNumber - 1) * consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&appies).Error; err != nil {
		return
	}
	return
}
