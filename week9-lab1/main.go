package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Book struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	ISBN       string    `json:"isbn"`
	Year       int       `json:"year"`
	Price      float64   `json:"price"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func initDB() {
	var err error
	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "bookstore")
	user := getEnv("DB_USER", "bookstore_user")
	password := getEnv("DB_PASSWORD", "your_strong_password")
	port := getEnv("DB_PORT", "5432")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	db, err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("Failed to open database.")
	}

	//กำหนดจำนวน Connection สูงสุด
	db.SetMaxOpenConns(25)

	// กำหนดจำนวน Idle connection สูงสุด
	db.SetMaxIdleConns(20)

	// กำหนดอายุของ Connection
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to Ping.", err)

	}

	log.Println("Successfully~~~")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getHealth(c *gin.Context) {
	err := db.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Unhealthy", "error": err})
		return
	}
	c.JSON(200, gin.H{"message": "healthy"})
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	var book Book

	// QueryRow ใช้เมื่อคาดว่าจะได้ผลลัพธ์ 0 หรือ 1 แถว
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func getAllBooks(c *gin.Context) {
	var rows *sql.Rows
	var err error
	YearQ := c.Query("year")
	if YearQ == "" {
		rows, err = db.Query("SELECT id, title, author, isbn, year, price, created_at, updated_at FROM books")
	} else {
		rows, err = db.Query("SELECT id, title, author, isbn, year, price, created_at, updated_at FROM books WHERE year = $1", YearQ)
	}
	// ลูกค้าถาม "มีหนังสืออะไรบ้าง"

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close() // ต้องปิด rows เสมอ เพื่อคืน Connection กลับ pool

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Year, &book.Price, &book.Created_At, &book.Updated_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		books = append(books, book)
	}
	if books == nil {
		books = []Book{}
	}
	c.JSON(http.StatusOK, books)

}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ใช้ RETURNING เพื่อดึงค่าที่ database generate (id, timestamps) eiei
	var id int
	var created_At, updated_At time.Time

	err := db.QueryRow(
		`INSERT INTO books (title, author, isbn, year, price)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id, created_at, updated_at`,
		newBook.Title, newBook.Author, newBook.ISBN, newBook.Year, newBook.Price,
	).Scan(&id, &created_At, &updated_At)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = id
	newBook.Created_At = created_At
	newBook.Updated_At = updated_At

	c.JSON(http.StatusCreated, newBook) // ใช้ 201 Created
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updateBook Book
	var ID int
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedAt time.Time
	err := db.QueryRow(
		`UPDATE books
         SET title = $1, author = $2, isbn = $3, year = $4, price = $5
         WHERE id = $6
         RETURNING id, updated_at`,
		updateBook.Title, updateBook.Author, updateBook.ISBN,
		updateBook.Year, updateBook.Price, id,
	).Scan(&ID, &updatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	updateBook.ID = ID
	updateBook.Updated_At = updatedAt
	c.JSON(http.StatusOK, updateBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func main() {
	initDB()
	defer db.Close() //Clear resource, when you finish.

	r := gin.Default()
	r.GET("/health", getHealth)

	api := r.Group("/api/v1")
	{
		api.GET("/books", getAllBooks)
		api.GET("/books/:id", getBook)
		api.POST("/books", createBook)
		api.PUT("/books/:id", updateBook)
		api.DELETE("/books/:id", deleteBook)
	}

	r.Run(":8080")
}
