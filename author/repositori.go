package author

import (
	"go_api_buku/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Author, error)
	FindById(ID int) (entity.Author, error)
	Create(author entity.Author) (entity.Author, error)
	Delete(author entity.Author) (entity.Author, error)
	Update(author entity.Author) (entity.Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Author, error) {
	var authors []entity.Author

	err := r.db.Find(&authors).Error

	return authors, err
}

func (r *repository) FindById(ID int) (entity.Author, error) {
	var author entity.Author

	err := r.db.First(&author, ID).Error

	return author, err
}

func (r *repository) Create(author entity.Author) (entity.Author, error) {
	err := r.db.Create(&author).Error
	return author, err
}

func (r *repository) Delete(author entity.Author) (entity.Author, error) {
	err := r.db.Delete(&author).Error

	return author, err
}

func (r *repository) Update(author entity.Author) (entity.Author, error) {

	err := r.db.Save(&author).Error
	return author, err
}
