package entity

type UserRole struct {
	UserId uint   `gorm:"primary_key;not null"`
	Role   string `gorm:"type:varchar(55);not null"`
}
