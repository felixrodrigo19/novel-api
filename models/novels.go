package models

import "gorm.io/gorm"

// Novel represents a novel entity
// swagger:model
type Novel struct {
	gorm.Model
	Id          int       `gorm:"primaryKey;autoIncrement" json:"id"`     // ID of the novel
	Title       string    `gorm:"unique;not null" json:"title"`           // Title of the novel
	Description string    `gorm:"not null" json:"description"`            // Description of the novel
	Language    string    `gorm:"not null" json:"language"`               // Language of the novel
	Type        string    `gorm:"not null" json:"type"`                   // Type of the novel
	Genres      []*Genre  `gorm:"many2many:novel_genres" json:"genres"`   // Genres associated with the novel
	Authors     []*Author `gorm:"many2many:novel_authors" json:"authors"` // Authors associated with the novel
	Year        int       `gorm:"not null" json:"year"`                   // Year of publication
}

// Genre represents a genre entity
// swagger:model
type Genre struct {
	gorm.Model
	Id     int      `gorm:"primaryKey;autoIncrement" json:"id"`   // ID of the genre
	Name   string   `gorm:"unique;not null" json:"name"`          // Name of the genre
	Novels []*Novel `gorm:"many2many:novel_genres" json:"novels"` // Novels associated with the genre
}

// Author represents an author entity
// swagger:model
type Author struct {
	gorm.Model
	Id     int      `gorm:"primaryKey;autoIncrement" json:"id"`    // ID of the author
	Name   string   `gorm:"unique;not null" json:"name"`           // Name of the author
	Novels []*Novel `gorm:"many2many:novel_authors" json:"novels"` // Novels associated with the author
}
