package models

type comment struct {
        Content    string      `gorm:"column:content"`
        CreateTime int         `gorm:"column:create_time"`
        ID         int         `gorm:"column:id;primary_key"`
        ImgURL     string      `gorm:"column:img_url"`
        MienID     int         `gorm:"column:mien_id"`
        Status     int         `gorm:"column:status"`
        Type       int         `gorm:"column:type"`
}

// TableName sets the insert table name for this struct type
func (c *comment) TableName() string {
        return "party_comment"
}