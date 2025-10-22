package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	_ "week13-lab6/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
)

// ===================== Response Types =====================
type ErrorResponse struct {
	Message string `json:"message"`
}

// ===================== Book Model =====================
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	Year      int       `json:"year"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ===================== Auth Models =====================
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // ไม่ส่งไปใน JSON
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         UserInfo `json:"user"`
}

type UserInfo struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ===================== JWT Claims =====================
type CustomClaims struct {
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB
var jwtSecret = []byte("my-super-secret-key-change-in-production-2024")

// ===================== Password Hashing Functions =====================
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// ===================== JWT Functions =====================
func generateAccessToken(userID int, username string, roles []string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func generateRefreshToken(userID int, username string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		Roles:    []string{}, // Refresh token ไม่ต้องเก็บ roles
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func verifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ===================== Database Helper Functions =====================
func getUserRoles(userID int) ([]string, error) {
	query := `
		SELECT r.name
		FROM roles r
		JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = $1
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func checkUserPermission(userID int, permission string) bool {
	query := `
		SELECT COUNT(*)
		FROM permissions p
		JOIN role_permissions rp ON p.id = rp.permission_id
		JOIN user_roles ur ON rp.role_id = ur.role_id
		WHERE ur.user_id = $1 AND p.name = $2
	`

	var count int
	err := db.QueryRow(query, userID, permission).Scan(&count)
	if err != nil {
		log.Printf("Error checking permission: %v", err)
		return false
	}

	return count > 0
}

func storeRefreshToken(userID int, token string, expiresAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, userID, token, expiresAt)
	return err
}

func revokeRefreshToken(token string) error {
	query := `
		UPDATE refresh_tokens
		SET revoked_at = NOW()
		WHERE token = $1 AND revoked_at IS NULL
	`
	_, err := db.Exec(query, token)
	return err
}

func isRefreshTokenValid(token string) (int, bool) {
	query := `
		SELECT user_id
		FROM refresh_tokens
		WHERE token = $1
		AND expires_at > NOW()
		AND revoked_at IS NULL
	`

	var userID int
	err := db.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return 0, false
	}

	return userID, true
}

func logAudit(userID int, action, resource string, resourceID interface{}, details map[string]interface{}, c *gin.Context) {
	detailsJSON, _ := json.Marshal(details)

	query := `
		INSERT INTO audit_logs
		(user_id, action, resource, resource_id, details, ip_address, user_agent)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	var resourceIDStr string
	if resourceID != nil {
		resourceIDStr = fmt.Sprintf("%v", resourceID)
	}

	db.Exec(query,
		userID,
		action,
		resource,
		resourceIDStr,
		detailsJSON,
		c.ClientIP(),
		c.GetHeader("User-Agent"),
	)
}

func initDB() {
	var err error

	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	// fmt.Println(conSt)
	db, err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("failed to open database")
	}

	// กำหนดจำนวน Connection สูงสุด
	db.SetMaxOpenConns(25)

	// กำหนดจำนวน Idle connection สูงสุด
	db.SetMaxIdleConns(25)

	// กำหนดอายุของ Connection
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	log.Println("successfully connected to database")
}

// ===================== Authentication Endpoints =====================
func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ดึงข้อมูล user จาก database
	var user User
	query := `
		SELECT id, username, email, password_hash, is_active
		FROM users
		WHERE username = $1
	`

	err := db.QueryRow(query, req.Username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.IsActive,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// ตรวจสอบว่า user active หรือไม่
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account is disabled"})
		return
	}

	// ตรวจสอบ password
	if err := verifyPassword(user.PasswordHash, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// ดึง roles ของ user
	roles, err := getUserRoles(user.ID)
	if err != nil {
		log.Printf("Error getting roles: %v", err)
		roles = []string{} // ถ้าดึงไม่ได้ให้เป็น empty array
	}

	// สร้าง tokens
	accessToken, err := generateAccessToken(user.ID, user.Username, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	refreshToken, err := generateRefreshToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	// บันทึก refresh token ในฐานข้อมูล
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	if err := storeRefreshToken(user.ID, refreshToken, expiresAt); err != nil {
		log.Printf("Error storing refresh token: %v", err)
		// ไม่ return error เพราะ token ยังใช้ได้
	}

	// อัพเดท last_login
	db.Exec("UPDATE users SET last_login = NOW() WHERE id = $1", user.ID)

	// Log audit
	logAudit(user.ID, "login", "auth", nil, gin.H{
		"username": user.Username,
	}, c)

	// ส่ง response
	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Roles:    roles,
		},
	})
}

func refreshTokenHandler(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ตรวจสอบ refresh token
	userID, valid := isRefreshTokenValid(req.RefreshToken)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	// ดึงข้อมูล user
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = $1", userID).Scan(&username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// ดึง roles
	roles, err := getUserRoles(userID)
	if err != nil {
		roles = []string{}
	}

	// สร้าง access token ใหม่
	accessToken, err := generateAccessToken(userID, username, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func logout(c *gin.Context) {
	// ดึง refresh token จาก request
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Revoke refresh token
	if err := revokeRefreshToken(req.RefreshToken); err != nil {
		log.Printf("Error revoking token: %v", err)
	}

	// Log audit (ถ้ามี user_id ใน context)
	if userID, exists := c.Get("user_id"); exists {
		logAudit(userID.(int), "logout", "auth", nil, nil, c)
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

// ===================== Middleware =====================
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง token จาก Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		// ตรวจสอบ format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Verify token
		claims, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// เก็บข้อมูล user ใน context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

func requirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// ตรวจสอบ permission
		hasPermission := checkUserPermission(userID.(int), permission)
		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"error":    "insufficient permissions",
				"required": permission,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ===================== Book Handlers =====================
// @Summary Get all books
// @Description Get details of books
// @Tags Books
// @Produce  json
// @Success 200  {array}  Book
// @Failure 500  {object}  ErrorResponse
// @Router  /books [get]
func getAllBooks(c *gin.Context) {
	var rows *sql.Rows
	var err error
	// ลูกค้าถาม "มีหนังสืออะไรบ้าง"
	rows, err = db.Query("SELECT id, title, author, isbn, year, price, created_at, updated_at FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close() // ต้องปิด rows เสมอ เพื่อคืน Connection กลับ pool

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Year, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			// handle error
		}
		books = append(books, book)
	}
	if books == nil {
		books = []Book{}
	}

	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	var book Book

	// QueryRow ใช้เมื่อคาดว่าจะได้ผลลัพธ์ 0 หรือ 1 แถว
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ใช้ RETURNING เพื่อดึงค่าที่ database generate (id, timestamps)
	var id int
	var createdAt, updatedAt time.Time

	err := db.QueryRow(
		`INSERT INTO books (title, author, isbn, year, price)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id, created_at, updated_at`,
		newBook.Title, newBook.Author, newBook.ISBN, newBook.Year, newBook.Price,
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = id
	newBook.CreatedAt = createdAt
	newBook.UpdatedAt = updatedAt

	// Log audit
	userID := c.GetInt("user_id")
	logAudit(userID, "create", "books", newBook.ID, gin.H{
		"title":  newBook.Title,
		"author": newBook.Author,
		"isbn":   newBook.ISBN,
	}, c)

	c.JSON(http.StatusCreated, newBook) // ใช้ 201 Created
}

func updateBook(c *gin.Context) {
	var ID int

	id := c.Param("id")
	var updateBook Book

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
	updateBook.UpdatedAt = updatedAt

	// Log audit
	userID := c.GetInt("user_id")
	logAudit(userID, "update", "books", updateBook.ID, gin.H{
		"title":  updateBook.Title,
		"author": updateBook.Author,
	}, c)

	c.JSON(http.StatusOK, updateBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	// Log audit
	userID := c.GetInt("user_id")
	logAudit(userID, "delete", "books", id, nil, c)

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

// @title           Bookstore API with Authentication
// @version         2.0
// @description     Bookstore API with JWT Authentication and RBAC Authorization
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())

	// ===================== Public Endpoints =====================
	// Swagger documentation
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint (for Docker healthcheck)
	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	// ===================== Authentication Endpoints =====================
	auth := r.Group("/auth")
	{
		auth.POST("/login", login)                 // Login และรับ tokens
		auth.POST("/refresh", refreshTokenHandler) // Refresh access token
		auth.POST("/logout", logout)               // Logout และ revoke token
	}

	// ===================== Protected API Endpoints =====================
	api := r.Group("/api/v1")
	api.Use(authMiddleware()) // ทุก endpoint ต้อง authenticate
	{
		// Books endpoints with permission checks
		api.GET("/books",
			requirePermission("books:read"),
			getAllBooks)

		api.GET("/books/:id",
			requirePermission("books:read"),
			getBook)

		api.POST("/books",
			requirePermission("books:create"),
			createBook)

		api.PUT("/books/:id",
			requirePermission("books:update"),
			updateBook)

		api.DELETE("/books/:id",
			requirePermission("books:delete"),
			deleteBook)
	}

	r.Run(":8080")
}
