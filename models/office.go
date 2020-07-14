package models

import (
	orm "danjian/database"
)

type Office struct {
	ID  int64 `json:"id"`
	UserId  int64 `json:"user_id"`
	Branch  int64 `json:"branch"`
	Role  int64 `json:"role"`
	IsSecretary  int64 `json:"is_secretary"`
	IsCommittee  int64 `json:"is_committee"`
	IsMember  int64 `json:"is_member"`
	CreateTime  int64 `json:"create_time"`
}

var Offices []Office

//列表
func (office *Office) List(page int)(offices []Office, err error) {
	if err = orm.Eloquent.Offset((page-1)*2).Limit(2).Find(&offices).Error; err != nil {
		return 
	}
	return
}