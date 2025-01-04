package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book 구조체 (임시 데이터 저장용)
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// 주석 추가 ---
// 메모리에 저장할 예시 데이터
var books = []Book{
	{ID: 1, Title: "Gin Framework", Author: "Gin Devs"},
	{ID: 2, Title: "Go Programming", Author: "John Doe"},
}

func main() {
	r := gin.Default()

	// Create (POST /books)
	r.POST("/books", func(c *gin.Context) {
		var newBook Book
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 임시 ID 할당 (실제 환경에서는 DB 에서 auto increment 등 사용)
		newBook.ID = len(books) + 1
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	})

	// feat/log 브랜치에서 추가된 코드
	// Read (GET /books) - 전체 조회
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// Read (GET /books/:id) - 단일 조회
	r.GET("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for _, book := range books {
			if book.ID == id {
				c.JSON(http.StatusOK, book)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Update (PUT /books/:id)
	r.PUT("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedBook Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, book := range books {
			if book.ID == id {
				books[i].Title = updatedBook.Title
				books[i].Author = updatedBook.Author
				c.JSON(http.StatusOK, books[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Delete (DELETE /books/:id)
	r.DELETE("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for i, book := range books {
			if book.ID == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// 서버 실행
	r.Run() // 기본 포트 : 8080
}
