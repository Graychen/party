package models

type dynamic struct {
        Address    string   `gorm:"column:address"`
        Content    string   `gorm:"column:content"`
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        Joined     string   `gorm:"column:joined"`
        Name       string   `gorm:"column:name"`
        Record     string   `gorm:"column:record"`
        Status     int      `gorm:"column:status"`
        Time       int      `gorm:"column:time"`
}

// TableName sets the insert table name for this struct type
func (d *dynamic) TableName() string {
        return "party_dynamic"
}