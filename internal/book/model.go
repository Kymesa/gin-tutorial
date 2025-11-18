package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" validate:"required,min=3,max=50"`
}
