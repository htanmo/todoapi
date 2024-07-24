package models

type Todo struct {
	ID          int    `json:"id" gorm:"primary key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
