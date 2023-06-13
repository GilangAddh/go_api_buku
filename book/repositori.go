package book

import (
	"go_api_buku/entity"

	"gorm.io/gorm"
)

type Repository interface {
	// db.Raw("Select * FROM book_auths")
	FindAll() ([]entity.BookAuth, error)
	FindById(ID int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.BookAuth, error) {
	var books []entity.BookAuth

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (entity.Book, error) {
	var book entity.Book

	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *repository) Create(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Delete(book entity.Book) (entity.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}

func (r *repository) Update(book entity.Book) (entity.Book, error) {

	err := r.db.Save(&book).Error
	return book, err
}
