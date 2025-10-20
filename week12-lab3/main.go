package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// กำหนด Claims structure
type CustomClaims struct {
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

var secretKey = []byte("my-super-secret-key-change-in-production")

// Mock user database
var users = map[string]User{
	"alice": {
		ID:       1,
		Username: "alice",
		Password: "password123",
		Roles:    []string{"admin"},
	},
	"bob": {
		ID:       2,
		Username: "bob",
		Password: "password456",
		Roles:    []string{"user"},
	},
}

// ฟังก์ชันสร้าง JWT
func generateToken(userID int, username string, roles []string) (string, error) {
	// กำหนดเวลาหมดอายุ (24 ชั่วโมง)
	expirationTime := time.Now().Add(24 * time.Hour)

	// สร้าง claims
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

	// สร้าง token ด้วย claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token ด้วย secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ฟังก์ชันตรวจสอบ JWT
func verifyToken(tokenString string) (*CustomClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบ algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// ดึง claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// Login
func login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	// ตรวจสอบ credentials
	user, exists := users[credentials.Username]
	if !exists || user.Password != credentials.Password {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	// สร้าง JWT
	token, err := generateToken(user.ID, user.Username, user.Roles)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	// ส่ง token กลับ
	c.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"roles":    user.Roles,
		},
	})
}

// Middleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง token จาก header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		// Format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token
		claims, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// เก็บข้อมูล user
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

// Middleware สำหรับตรวจสอบ role
func requireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists {
			c.JSON(403, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		rolesList := roles.([]string)
		hasRole := false
		for _, role := range rolesList {
			if role == requiredRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(403, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Public routes
	r.POST("/login", login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			roles, _ := c.Get("roles")
			c.JSON(200, gin.H{
				"username": username,
				"roles":    roles,
			})
		})

		// Admin only route
		protected.GET("/admin", requireRole("admin"), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome admin!",
			})
		})
	}

	fmt.Println("Server running on :9999")
	fmt.Println("Try:")
	fmt.Println("  curl -X POST http://localhost:9999/login -H 'Content-Type: application/json' -d '{\"username\":\"alice\",\"password\":\"password123\"}'")
	r.Run(":9999")
}
