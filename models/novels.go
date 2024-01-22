package models

import "gorm.io/gorm"

type Novel struct {
	gorm.Model
	Id          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"unique;not null"`
	Description string    `gorm:"not null"`
	Language    string    `gorm:"not null"`
	Type        string    `gorm:"not null"`
	Genres      []*Genre  `gorm:"many2many:novel_genres"`
	Authors     []*Author `gorm:"many2many:novel_authors"`
	Year        int       `gorm:"not null"`
}

type Genre struct {
	gorm.Model
	Id     int      `gorm:"primaryKey;autoIncrement"`
	Name   string   `gorm:"unique;not null"`
	Novels []*Novel `gorm:"many2many:novel_genres"`
}

type Author struct {
	gorm.Model
	Id     int      `gorm:"primaryKey;autoIncrement"`
	Name   string   `gorm:"unique;not null"`
	Novels []*Novel `gorm:"many2many:novel_authors"`
}
