package models

type theoy struct {
        Content    string   `gorm:"column:content"`
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        Num        int      `gorm:"column:num"`
        Status     int      `gorm:"column:status"`
        Title      string   `gorm:"column:title"`
}

// TableName sets the insert table name for this struct type
func (t *theoy) TableName() string {
        return "party_theoy"
}