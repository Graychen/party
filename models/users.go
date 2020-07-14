package models

type users struct {
        ID       int    `gorm:"column:id;primary_key"`
        Name     string `gorm:"column:name"`
        Password string `gorm:"column:password"`
}

// TableName sets the insert table name for this struct type
func (u *users) TableName() string {
        return "party_users"
}