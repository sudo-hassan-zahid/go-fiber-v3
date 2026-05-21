package models

type Tokens struct {
	BaseModel
	UserID       uint   `json:"user_id" gorm:"not null"`
	RefreshToken string `json:"refresh_token" gorm:"type:varchar(255);not null"`
	User         User   `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
