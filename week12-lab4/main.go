package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT + Refresh Token + Token Blacklist
// 1. Login >> สร้าง JWT
// 2. เก็บ JWT ใน httpOnly cookie (ไม่ใช่ localStorage)
// 3. Backend verify JWT แต่ check blacklist ด้วย

// กำหนด Claims structure
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var secretKey = []byte("my-super-secret-key-change-in-production")

// Mock user database
var users = map[string]User{
	"alice": {
		ID:       1,
		Username: "alice",
		Password: "password123",
	},
	"bob": {
		ID:       2,
		Username: "bob",
		Password: "password456",
	},
}

// Token blacklist (in-memory)
var blacklist = struct {
	sync.RWMutex
	tokens map[string]bool
}{tokens: make(map[string]bool)}

// Refresh token storage (in-memory)
var refreshTokenStore = struct {
	sync.RWMutex
	tokens map[int]string // userID -> refreshToken
}{tokens: make(map[int]string)}

// ฟังก์ชันสร้าง JWT
func generateToken(user User, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	claims := &CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ฟังก์ชันตรวจสอบ JWT
func verifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// Blacklist functions
func isBlacklisted(token string) bool {
	blacklist.RLock()
	defer blacklist.RUnlock()
	return blacklist.tokens[token]
}

func addToBlacklist(token string) {
	blacklist.Lock()
	defer blacklist.Unlock()
	blacklist.tokens[token] = true
}

// Refresh token functions
func storeRefreshToken(userID int, token string) {
	refreshTokenStore.Lock()
	defer refreshTokenStore.Unlock()
	refreshTokenStore.tokens[userID] = token
}

func getRefreshToken(userID int) (string, bool) {
	refreshTokenStore.RLock()
	defer refreshTokenStore.RUnlock()
	token, exists := refreshTokenStore.tokens[userID]
	return token, exists
}

func deleteRefreshToken(userID int) {
	refreshTokenStore.Lock()
	defer refreshTokenStore.Unlock()
	delete(refreshTokenStore.tokens, userID)
}

// Login handler
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

	// สร้าง Access Token (อายุสั้น: 15 นาที)
	accessToken, err := generateToken(user, 15*time.Minute)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}

	// สร้าง Refresh Token (อายุยาว: 7 วัน)
	refreshToken, err := generateToken(user, 7*24*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate refresh token"})
		return
	}

	// เก็บ access token ใน httpOnly cookie
	// maxAge: 900 seconds = 15 minutes
	// secure=false เพราะใช้ HTTP (production ควรเป็น true สำหรับ HTTPS)
	// httpOnly=true >> JavaScript ไม่สามารถอ่านได้ (ป้องกัน XSS)
	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)

	// เก็บ refresh token ใน httpOnly cookie
	// maxAge: 604800 seconds = 7 days
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "", false, true)

	// เก็บ refresh token ใน store
	storeRefreshToken(user.ID, refreshToken)

	c.JSON(200, gin.H{
		"message": "logged in",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

// Auth middleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง token จาก cookie
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized - no token"})
			c.Abort()
			return
		}

		// Verify JWT
		claims, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Check blacklist
		if isBlacklisted(tokenString) {
			c.JSON(401, gin.H{"error": "token revoked"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

// Refresh token handler
func refresh(c *gin.Context) {
	// ดึง refresh token จาก cookie
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "no refresh token"})
		return
	}

	// Verify refresh token
	claims, err := verifyToken(refreshTokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid refresh token"})
		return
	}

	// ตรวจสอบว่า refresh token ตรงกับที่เก็บไว้ใน store หรือไม่
	storedToken, exists := getRefreshToken(claims.UserID)
	if !exists || storedToken != refreshTokenString {
		c.JSON(401, gin.H{"error": "refresh token not found"})
		return
	}

	// สร้าง access token ใหม่
	user := User{
		ID:       claims.UserID,
		Username: claims.Username,
	}

	newAccessToken, err := generateToken(user, 15*time.Minute)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate new token"})
		return
	}

	// ส่ง access token ใหม่
	c.SetCookie("access_token", newAccessToken, 900, "/", "", false, true)

	c.JSON(200, gin.H{"message": "token refreshed"})
}

// Logout handler
func logout(c *gin.Context) {
	// ดึง access token
	accessToken, _ := c.Cookie("access_token")
	if accessToken != "" {
		// เพิ่มเข้า blacklist
		addToBlacklist(accessToken)
	}

	// ดึง user_id เพื่อลบ refresh token
	userID, exists := c.Get("user_id")
	if exists {
		deleteRefreshToken(userID.(int))
	}

	// ลบ cookies
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(200, gin.H{"message": "logged out"})
}

func main() {
	r := gin.Default()

	// Public routes
	r.POST("/login", login)
	r.POST("/refresh", refresh)

	// Protected routes
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"user_id":  userID,
				"username": username,
			})
		})

		protected.POST("/logout", logout)
	}

	fmt.Println("Server running on :9999")
	r.Run(":9999")
}
