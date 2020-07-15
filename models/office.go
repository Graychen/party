package models

import (
	orm "danjian/database"
)

type Office struct {
	ID  int64 `gorm:"column:id"`
	UserId  int64 `gorm:"column:user_id"`
	Branch  int64 `gorm:"column:branch"`
	Role  int64 `gorm:"column:role"`
	IsSecretary  int64 `gorm:"column:is_secretary"`
	IsCommittee  int64 `gorm:"column:is_committee"`
	IsMember  int64 `gorm:"column:is_member"`
	CreateTime  int64 `gorm:"column:create_time"`
}

var Offices []Office

//列表
func (office *Office) List(page int)(offices []Office, err error) {
	if err = orm.Eloquent.Offset((page-1)*2).Limit(2).Find(&offices).Error; err != nil {
		return 
	}
	return
}