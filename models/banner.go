package models

import (
	orm "danjian/database"
	    "danjian/consts"
)

type Banner struct {
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        ImgURL     string   `gorm:"column:img_url"`
        Status     int      `gorm:"column:status"`
        UserID     int      `gorm:"column:user_id"`
}

//列表
func (banner *Banner) List(page int)(banners []Banner, err error) {
	if err = orm.Eloquent.Offset((page-1)*consts.BANNER_SIZE).Limit(consts.BANNER_SIZE).Find(&banners).Error; err != nil {
		return 
	}
	return
}