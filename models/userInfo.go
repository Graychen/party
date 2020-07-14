package models

type user_info struct {
        CreateTime int      `gorm:"column:create_time"`
        Email      string   `gorm:"column:email"`
        HiredDate  int      `gorm:"column:hiredDate"`
        ID         int      `gorm:"column:id;primary_key"`
        IsSenior   int      `gorm:"column:is_senior"`
        Jobnumber  string   `gorm:"column:jobnumber"`
        Mobile     int      `gorm:"column:mobile"`
        Tel        int      `gorm:"column:tel"`
        UserID     int      `gorm:"column:user_id"`
        WorkPlace  string   `gorm:"column:work_place"`
}

// TableName sets the insert table name for this struct type
func (u *user_info) TableName() string {
        return "party_user_info"
}