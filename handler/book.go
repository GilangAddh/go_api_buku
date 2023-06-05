package handler

import (
	// "encoding/json"

	"fmt"
	"go_api_buku/book"
	"go_api_buku/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	b, err := h.bookService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	idb, _ := strconv.Atoi(id)

	b, err := h.bookService.Delete(int(idb))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToResponse(b)

	deleteMessage := "Sucessfull Deleting book " + bookResponse.Title

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var BookUpdate book.BookUpdate

	err := c.ShouldBindJSON(&BookUpdate)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	id := c.Param("id")
	idb, err := strconv.Atoi(id)
	book, err := h.bookService.Update(idb, BookUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})

}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func convertToResponse(b entity.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		AuthorID:    b.AuthorID,
	}
}