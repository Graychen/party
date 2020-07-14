package models

type banner struct {
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        ImgURL     string   `gorm:"column:img_url"`
        Status     int      `gorm:"column:status"`
        UserID     int      `gorm:"column:user_id"`
}

// TableName sets the insert table name for this struct type
func (b *banner) TableName() string {
        return "party_banner"
}