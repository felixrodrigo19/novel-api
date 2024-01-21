package models

import "gorm.io/gorm"

type Novel struct {
	gorm.Model
	Id          int      `gorm:"primaryKey,autoIncrement" json:"id"`
	Title       string   `gorm:"unique,not null" json:"title"`
	Description string   `gorm:"not null" json:"description"`
	Language    string   `gorm:"not null" json:"language"`
	Type        string   `gorm:"not null" json:"type"`
	Genres      []Genre  `gorm:"many2many:novel_genres" json:"genres"`
	Authors     []Author `gorm:"many2many:novel_authors" json:"authors"`
	Year        int      `gorm:"not null" json:"year"`
}

type Genre struct {
	gorm.Model
	Id     int     `gorm:"primaryKey,autoIncrement" json:"genre-id"`
	Name   string  `gorm:"unique,not null" json:"name"`
	Novels []Novel `gorm:"many2many:novel_genres" json:"novels"`
}

type Author struct {
	gorm.Model
	Id     int     `gorm:"primaryKey,autoIncrement" json:"author-id"`
	Name   string  `gorm:"unique,not null" json:"name"`
	Novels []Novel `gorm:"many2many:novel_authors" json:"novels"`
}
