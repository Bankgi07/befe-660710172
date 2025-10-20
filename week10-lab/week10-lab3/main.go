package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "week10-lab3/docs"
)

// ===== Swagger Base Info =====
//
// @title           Simple API Example
// @version         1.0
// @description     This is a simple example of using Gin with Swagger.
// @host            localhost:8080
// @BasePath        /api/v1

type ErrorResponse struct {
	Message string `json:"message"`
}

var db *sql.DB

// ใช้ตอบกลับ
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

// ใช้ตอนสร้าง
type CreateBookRequest struct {
	Title  string  `json:"title"  binding:"required"`
	Author string  `json:"author" binding:"required"`
	ISBN   string  `json:"isbn"   binding:"required"`
	Year   int     `json:"year"   binding:"required"`
	Price  float64 `json:"price"  binding:"required"`
}

// ใช้ตอนแก้ไข
type UpdateBookRequest struct {
	Title  string  `json:"title"  binding:"required"`
	Author string  `json:"author" binding:"required"`
	ISBN   string  `json:"isbn"   binding:"required"`
	Year   int     `json:"year"   binding:"required"`
	Price  float64 `json:"price"  binding:"required"`
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func initDB() {
	var err error

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "postgres")
	name := getEnv("DB_NAME", "mydatabase")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		// ช่วง dev ให้ไปต่อได้เพื่อเปิด Swagger/เอกสาร
		log.Println("[WARN] Failed to Ping DB:", err)
	} else {
		log.Println("[DB] Connected successfully")
	}
}

// ===== Health =====

// @Summary Health check
// @Tags    Health
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 503 {object} ErrorResponse
// @Router  /health [get]
func getHealth(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Message: "DB not initialized"})
		return
	}
	if err := db.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "healthy"})
}

// ===== Handlers =====

// @Summary      Get all books (optional filter by year)
// @Description  Get all books; filter by year with query ?year=YYYY
// @Tags         Books
// @Produce      json
// @Param        year  query     int  false  "Filter by year"
// @Success      200   {array}   Book
// @Failure      500   {object}  ErrorResponse
// @Router       /books [get]
func getAllBooks(c *gin.Context) {
	var rows *sql.Rows
	var err error

	yearQ := c.Query("year")
	if yearQ == "" {
		rows, err = db.Query(`SELECT id, title, author, isbn, year, price, created_at, updated_at FROM books`)
	} else {
		rows, err = db.Query(`SELECT id, title, author, isbn, year, price, created_at, updated_at FROM books WHERE year = $1`, yearQ)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Year, &b.Price, &b.Created_At, &b.Updated_At); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
			return
		}
		books = append(books, b)
	}
	if books == nil {
		books = []Book{}
	}
	c.JSON(http.StatusOK, books)
}

// @Summary      Get new books
// @Description  Latest 5 books ordered by created_at DESC
// @Tags         Books
// @Produce      json
// @Success      200  {array}   Book
// @Failure      500  {object}  ErrorResponse
// @Router       /books/new [get]
func getNewBooks(c *gin.Context) {
	rows, err := db.Query(`
		SELECT id, title, author, isbn, year, price, created_at, updated_at
		FROM books ORDER BY created_at DESC LIMIT 5
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Year, &b.Price, &b.Created_At, &b.Updated_At); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
			return
		}
		books = append(books, b)
	}
	if books == nil {
		books = []Book{}
	}
	c.JSON(http.StatusOK, books)
}

// @Summary      Get book by ID
// @Description  Get one book by its ID
// @Tags         Books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  Book
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /books/{id} [get]
func getBook(c *gin.Context) {
	id := c.Param("id")
	var b Book
	err := db.QueryRow(`
		SELECT id, title, author, isbn, year, price, created_at, updated_at
		FROM books WHERE id = $1
	`, id).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Year, &b.Price, &b.Created_At, &b.Updated_At)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

// @Summary      Create a new book
// @Description  Create a new book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        book  body      CreateBookRequest  true  "Book payload"
// @Success      201   {object}  Book
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /books [post]
func createBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	var b Book
	err := db.QueryRow(`
		INSERT INTO books (title, author, isbn, year, price)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, author, isbn, year, price, created_at, updated_at
	`, req.Title, req.Author, req.ISBN, req.Year, req.Price).Scan(
		&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Year, &b.Price, &b.Created_At, &b.Updated_At,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, b)
}

// @Summary      Update a book
// @Description  Update a book by ID
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id    path      int                true  "Book ID"
// @Param        book  body      UpdateBookRequest  true  "Book payload"
// @Success      200   {object}  Book
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /books/{id} [put]
func updateBook(c *gin.Context) {
	id := c.Param("id")

	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	var b Book
	err := db.QueryRow(`
		UPDATE books
		SET title = $1, author = $2, isbn = $3, year = $4, price = $5
		WHERE id = $6
		RETURNING id, title, author, isbn, year, price, created_at, updated_at
	`, req.Title, req.Author, req.ISBN, req.Year, req.Price, id).Scan(
		&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Year, &b.Price, &b.Created_At, &b.Updated_At,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

// @Summary      Delete a book
// @Description  Delete a book by ID
// @Tags         Books
// @Produce      json
// @Param        id  path      int  true  "Book ID"
// @Success      200 {object}  map[string]string
// @Failure      404 {object}  ErrorResponse
// @Failure      500 {object}  ErrorResponse
// @Router       /books/{id} [delete]
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	res, err := db.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func main() {
	initDB()
	if db != nil {
		defer db.Close()
	}

	r := gin.Default()
	// CORS เปิดกว้างช่วง dev
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	r.Use(cors.New(c))

	// Health
	r.GET("/health", getHealth)

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1
	api := r.Group("/api/v1")
	{
		api.GET("/books", getAllBooks)
		api.GET("/books/new", getNewBooks)
		api.GET("/books/:id", getBook)
		api.POST("/books", createBook)
		api.PUT("/books/:id", updateBook)
		api.DELETE("/books/:id", deleteBook)
	}

	// bind 0.0.0.0 ให้ container/port-forward เห็นแน่ๆ
	r.Run("0.0.0.0:8080")
	r.GET("/", func(c *gin.Context) { c.Redirect(302, "/docs/index.html") })
	r.GET("/index.html", func(c *gin.Context) { c.Redirect(302, "/docs/index.html") })
}
