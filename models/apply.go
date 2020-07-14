package models

type apply struct {
        ActivityID int      `gorm:"column:activity_id"`
        CreateTime int      `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        Status     int      `gorm:"column:status"`
        UserID     int      `gorm:"column:user_id"`
}

// TableName sets the insert table name for this struct type
func (t *apply) TableName() string {
        return "party_apply"
}