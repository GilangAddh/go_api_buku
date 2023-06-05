package author

import (
	"fmt"
	"go_api_buku/entity"
)

type Service interface {
	FindAll() ([]entity.Author, error)
	FindById(ID int) (entity.Author, error)
	Create(author AuthorRequest) (entity.Author, error)
	Delete(ID int) (entity.Author, error)
	Update(ID int, author AuthorRequest) (entity.Author, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.Author, error) {
	authors, err := s.repository.FindAll()
	return authors, err
}

func (s *service) FindById(ID int) (entity.Author, error) {
	author, err := s.repository.FindById(ID)
	return author, err
}

func (s *service) Create(authorRequest AuthorRequest) (entity.Author, error) {
	author := entity.Author{
		Name: authorRequest.Name,
	}
	newAuthor, err := s.repository.Create(author)
	return newAuthor, err
}

func (s *service) Delete(ID int) (entity.Author, error) {
	author, err := s.repository.FindById(ID)
	delAuthor, err := s.repository.Delete(author)

	return delAuthor, err
}

func (s *service) Update(ID int, authorRequest AuthorRequest) (entity.Author, error) {
	author, err := s.repository.FindById(ID)

	author.Name = authorRequest.Name

	updateAuthor, err := s.repository.Update(author)
	fmt.Println(updateAuthor)
	return updateAuthor, err
}
