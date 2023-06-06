package main

import (
	// "encoding/json"
	// "fmt"
	// "net/http"
	"log"

	"github.com/gin-gonic/gin"

	//"github.com/go-playground/validator/v10"
	"go_api_buku/author"
	"go_api_buku/book"
	"go_api_buku/entity"
	"go_api_buku/handler"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// mysql
	dsn := "root:@tcp(127.0.0.1:3306)/buku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Postgre
	// dsn := "host=localhost user=postgres password=superadmin dbname=buku port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection Error")
	}
	// migrate
	db.AutoMigrate(&entity.Author{})
	db.AutoMigrate(&entity.Book{})
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	authorHandler := handler.NewAuthorHandler(author.NewService(author.NewRepository(db)))

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBook)

	v1.POST("/author", authorHandler.CreateAuthor)
	v1.GET("/author", authorHandler.GetAuthors)
	v1.GET("/author/:id", authorHandler.GetAuthor)
	v1.DELETE("/author/:id", authorHandler.DeleteAuthor)
	v1.PUT("/author/:id", authorHandler.UpdateAuthor)

	//Kalau kita mau 2 versi ketika ada update pada API
	// v2 := router.Group("/v2")

	// v2.GET("/", rootHandler)
	// v2.GET("books/:id/:title", booksHandler)
	// v2.GET("/query", queryHandler)
	// v2.POST("/books", postBooksHandler)

	//ganti port taro di parameter run
	router.Run(":8081")
}
