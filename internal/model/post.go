package model

import "time"

type Post struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique; not null" json:"username"`
	Email    string `gorm:"unique; not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	// Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
