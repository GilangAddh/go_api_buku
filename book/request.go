package book

import (
	"encoding/json"
)

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	AuthorID    json.Number `json:"AuthorID" binding:"required,number"`
}

type BookUpdate struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	AuthorID    json.Number `json:"AuthorID" binding:"required,number"`
}
