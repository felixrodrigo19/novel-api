package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Id   int    `json:"genre-id"`
	Name string `json:"genre-name"`
}

type Author struct {
	gorm.Model
	Id   int    `json:"author-id"`
	Name string `json:"author-name"`
}

type Novel struct {
	gorm.Model
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Type        string `json:"type"`
	Genre       `json:"genre"`
	Author      `json:"author"`
	Year        int64 `json:"year"`
}

var Novels []Novel
