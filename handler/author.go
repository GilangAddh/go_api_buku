package handler

import (
	"fmt"
	"go_api_buku/author"
	"go_api_buku/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/go-playground/validator/v10"
)

// "encoding/json"

type authorHandler struct {
	authorService author.Service
}

func NewAuthorHandler(authorService author.Service) *authorHandler {
	return &authorHandler{authorService}
}

func (h *authorHandler) GetAuthors(c *gin.Context) {
	CorsPoliciy(c)
	authors, err := h.authorService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var authorsResponse []author.AuthorResponse

	for _, a := range authors {
		authorResponse := convertAToResponse(a)
		authorsResponse = append(authorsResponse, authorResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": authorsResponse,
	})
}

func (h *authorHandler) GetAuthor(c *gin.Context) {
	CorsPoliciy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	a, err := h.authorService.FindById(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	authorResponse := convertAToResponse(a)

	c.JSON(http.StatusOK, gin.H{
		"data": authorResponse,
	})
}

func (h *authorHandler) DeleteAuthor(c *gin.Context) {
	CorsPoliciy(c)
	id := c.Param("id")
	idb, err := strconv.Atoi(id)

	a, err := h.authorService.Delete(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	authorResponse := convertAToResponse(a)

	deleteMessage := "Sucessfull Deleting author " + authorResponse.Name

	c.JSON(http.StatusOK, gin.H{
		"data": deleteMessage,
	})
}

func (h *authorHandler) CreateAuthor(c *gin.Context) {
	CorsPoliciy(c)
	var authorRequest author.AuthorRequest

	err := c.ShouldBindJSON(&authorRequest)
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

	author, err := h.authorService.Create(authorRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertAToResponse(author),
	})

}

func (h *authorHandler) UpdateAuthor(c *gin.Context) {
	CorsPoliciy(c)
	var authorRequest author.AuthorRequest

	err := c.ShouldBindJSON(&authorRequest)
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

	a, err := h.authorService.Update(idb, authorRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertAToResponse(a),
	})

}

func convertAToResponse(a entity.Author) author.AuthorResponse {
	return author.AuthorResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}
