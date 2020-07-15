package models

import (
	orm "danjian/database"
	    "danjian/consts"
)

type Activity struct {
	ID  int64 `json:"id"`
	Theme  string `json:"theme"`
	Time  int64 `json:"time"`
	Address  string `json:"address"`
	Content  string `json:"content"`
	Totol  int64 `json:"totol"`
	Status  int64 `json:"status"`
	CreateTime  int64 `json:"create_time"`
}

var Activities []Activity

//列表
func (activity *Activity) List(page int)(activities []Activity, err error) {
	if err = orm.Eloquent.Offset((page-1)*consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&activities).Error; err != nil {
		return 
	}
	return
}