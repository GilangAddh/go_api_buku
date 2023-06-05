package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	AuthorID    int
	Author      Author
}

type Author struct {
	ID   int
	Name string
}
