package models

type Register struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email"` 
	Password string `gorm:"not null" json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}