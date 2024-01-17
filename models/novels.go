package models

type Genre struct {
	Name string `json:"genre-name"`
}

type Author struct {
	Name string `json:"authon-name"`
}

type Novel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Type        string `json:"type"`
	Genre       `json:"novel-genre"`
	Author      `json:"novel-author"`
	Year        int64 `json:"year"`
}
