package models

import (
	orm "danjian/database"
)

type Office struct {
	ID  int `gorm:"column:id;primary_key"`
	UserId  int `gorm:"column:user_id"`
	Branch  int `gorm:"column:branch"`
	Role  int `gorm:"column:role"`
	IsSecretary  int `gorm:"column:is_secretary"`
	IsCommittee  int `gorm:"column:is_committee"`
	IsMember  int `gorm:"column:is_member"`
	CreateTime  int `gorm:"column:create_time"`
}

var Offices []Office

//列表
func (office *Office) List(page int)(offices []Office, err error) {
	if err = orm.Eloquent.Offset((page-1)*2).Limit(2).Find(&offices).Error; err != nil {
		return
	}
	return
}