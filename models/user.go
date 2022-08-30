package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"default:customer"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Phone     int       `json:"phone" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	PostCode  string    `json:"post" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
