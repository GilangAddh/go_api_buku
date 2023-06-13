package book

import (
	"fmt"
	"go_api_buku/entity"
)

type Service interface {
	FindAll() ([]entity.BookAuth, error)
	FindById(ID int) (entity.Book, error)
	Create(book BookRequest) (entity.Book, error)
	Delete(ID int) (entity.Book, error)
	Update(ID int, book BookUpdate) (entity.Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.BookAuth, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (entity.Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (entity.Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	authorID, _ := bookRequest.AuthorID.Int64()
	book := entity.Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Rating:      int(rating),
		Description: bookRequest.Description,
		AuthorID:    int(authorID),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Delete(ID int) (entity.Book, error) {
	book, err := s.repository.FindById(ID)
	delBook, err := s.repository.Delete(book)

	return delBook, err
}

func (s *service) Update(ID int, bookUpdate BookUpdate) (entity.Book, error) {
	book, err := s.repository.FindById(ID)

	price, _ := bookUpdate.Price.Int64()
	rating, _ := bookUpdate.Rating.Int64()
	authorID, _ := bookUpdate.AuthorID.Int64()

	book.Title = bookUpdate.Title
	book.Price = int(price)
	book.Rating = int(rating)
	book.Description = bookUpdate.Description
	book.AuthorID = int(authorID)

	updateBook, err := s.repository.Update(book)
	fmt.Println(updateBook)
	return updateBook, err
}
