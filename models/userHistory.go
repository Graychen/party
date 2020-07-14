package models

type user_history struct {
        CreateTime int `gorm:"column:create_time"`
        ID         int      `gorm:"column:id;primary_key"`
        URL        string   `gorm:"column:url"`
        UserID     int      `gorm:"column:user_id"`
}

// TableName sets the insert table name for this struct type
func (u *user_history) TableName() string {
        return "party_user_history"
}