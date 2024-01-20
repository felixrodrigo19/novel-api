package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Id   int    `gorm:"primaryKey,autoIncrement" json:"genre-id"`
	Name string `gorm:"unique,not null" json:"genre-name"`
}

type Author struct {
	gorm.Model
	Id   int    `gorm:"primaryKey,autoIncrement" json:"author-id"`
	Name string `gorm:"unique,not null" json:"author-name"`
}

type Novel struct {
	gorm.Model
	Id          int    `gorm:"primaryKey,autoIncrement" json:"id"`
	Title       string `gorm:"unique,not null" json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Type        string `json:"type"`
	Genre       `json:"genre"`
	Author      `json:"author"`
	Year        int64 `json:"year"`
}

var (
	Authors []Author
	Genres  []Genre
	Novels  []Novel
)
