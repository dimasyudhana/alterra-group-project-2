package entities

import (
	"time"

	"gorm.io/gorm"
)

type Core struct {
	ID       uint
	Title    string
	Year     string
	Author   string
	Contents string
	Image    string
	UserID   uint
}

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string         `json:"title" gorm:"type:varchar(225)"`
	Year      string         `json:"year" gorm:"type:varchar(4)"`
	Author    string         `json:"author" gorm:"type:varchar(225)"`
	Contents  string         `json:"contents"`
	Image     string         `json:"image" gorm:"type:blob"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UserID    uint           `json:"username" gorm:"type:varchar(15)"`
}
