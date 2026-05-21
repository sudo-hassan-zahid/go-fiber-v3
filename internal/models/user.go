package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;uniqueIndex"`
	Password  string `json:"password" gorm:"type:varchar(100);not null"`
}
