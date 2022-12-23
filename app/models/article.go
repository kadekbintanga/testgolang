package models

import(
	"time"
	"gorm.io/gorm"
)

type Article struct{
	gorm.Model
	ID			uint		`json:"id" gorm:"primary_key"`
	Author		string		`json:"author"`
	Title		string		`json:"title" gorm:"index:email_unique_index;unique"`
	Body		string		`json:"body"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
}