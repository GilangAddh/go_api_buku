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
	CorsPoliciy(c)
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponseView

	for _, b := range books {
		bookResponse := convertToResponseView(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	// mengembalikan data dalam bentuk Object -> array of object
	// c.JSON(http.StatusOK, gin.H{

	// 	"data": booksResponse,
	// })

	// mengembalikan data dalam bentuk array of object
	c.JSON(http.StatusOK, booksResponse)
}

func (h *bookHandler) GetBook(c *gin.Context) {
	CorsPoliciy(c)
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

	c.JSON(http.StatusOK, bookResponse)
}

func (h *bookHandler) DeleteHandler(c *gin.Context) {
	CorsPoliciy(c)
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
	CorsPoliciy(c)
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

	c.JSON(http.StatusOK, convertToResponse(book))

}

func (h *bookHandler) CreateBook(c *gin.Context) {
	CorsPoliciy(c)
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

	c.JSON(http.StatusOK, book)

}

func convertToResponseView(b entity.BookAuth) book.BookResponseView {
	return book.BookResponseView{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		AuthorID:    b.AuthorID,
		AuthorName:  b.AuthorName,
	}
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
