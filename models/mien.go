package models

import (
	orm "danjian/database"
	    "danjian/consts"
)

type Mien struct {
        Content    string      `gorm:"column:content"`
        CreateTime int         `gorm:"column:create_time"`
        ID         int         `gorm:"column:id;primary_key"`
        Num        int         `gorm:"column:num"`
        Status     int         `gorm:"column:status"`
        Title      string      `gorm:"column:title"`
        Type       int         `gorm:"column:type"`
}

// TableName sets the insert table name for this struct type
func (m *Mien) TableName() string {
        return "party_mien"
}

var Miens []Mien

//列表
func (mien *Mien) List(mienType int, page int)(miens []Mien, err error) {
	if mienType, isExist := mienType, isExist == true {
		orm.Eloquent.Where("type = ?", mienType)
	}
	if err = orm.Eloquent.Offset((page-1)*consts.PAGE_SIZE).Limit(consts.PAGE_SIZE).Find(&mien).Error; err != nil {
		return 
	}
	return
}