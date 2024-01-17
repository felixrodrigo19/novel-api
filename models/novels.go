package models

type Genre struct {
	Id   int    `json:"genre-id"`
	Name string `json:"genre-name"`
}

type Author struct {
	Id   int    `json:"author-id"`
	Name string `json:"author-name"`
}

type Novel struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Type        string `json:"type"`
	Genre       `json:"novel-genre"`
	Author      `json:"novel-author"`
	Year        int64 `json:"year"`
}
