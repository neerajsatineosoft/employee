package models

type Login struct {
	UUID     int    `gorm json:"user_id" form:"-"`
	Email    string `gorm:"size:50;not null" json:"email" form:"username"`
	Password string `gorm:"size:50;not null"json:"password" form:"password"`
}
