package main

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Server-side session storage
type SessionStore struct {
	sessions map[string]*SessionData
	mu       sync.RWMutex
}

type SessionData struct {
	UserID     int
	Username   string
	Roles      []string
	CreatedAt  time.Time
	LastAccess time.Time
}

var store = &SessionStore{
	sessions: make(map[string]*SessionData),
}

type User struct {
	ID       int
	Username string
	Roles    []string
}

func generateRandomID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Login
func login(c *gin.Context) {
	// For demo purposes, creating a mock user
	user := User{
		ID:       1,
		Username: "testuser",
		Roles:    []string{"admin"},
	}

	// สร้าง session ID
	sessionID := generateRandomID()

	// เก็บ session data
	store.mu.Lock()
	store.sessions[sessionID] = &SessionData{
		UserID:     user.ID,
		Username:   user.Username,
		Roles:      user.Roles,
		CreatedAt:  time.Now(),
		LastAccess: time.Now(),
	}
	store.mu.Unlock()

	// ส่ง cookie
	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)
	// httpOnly=true >> JavaScript access ไม่ได้ (ป้องกัน XSS)
	c.JSON(200, gin.H{"message": "logged in"})
}

// Middleware
func sessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// ดึง session data
		store.mu.RLock()
		session, exists := store.sessions[sessionID]
		store.mu.RUnlock()

		if !exists {
			c.JSON(401, gin.H{"error": "invalid session"})
			c.Abort()
			return
		}

		// Update last access
		store.mu.Lock()
		session.LastAccess = time.Now()
		store.mu.Unlock()

		// เก็บข้อมูล user ใน context
		c.Set("user_id", session.UserID)
		c.Set("username", session.Username)
		c.Set("roles", session.Roles)

		c.Next()
	}
}

// Logout
func logout(c *gin.Context) {
	sessionID, _ := c.Cookie("session_id")

	// ลบ session
	store.mu.Lock()
	delete(store.sessions, sessionID)
	store.mu.Unlock()

	// ลบ cookie
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"message": "logged out"})
}

func main() {
	r := gin.Default()

	r.POST("/login", login)
	r.POST("/logout", logout)

	// Protected routes
	protected := r.Group("/")
	protected.Use(sessionMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			roles, _ := c.Get("roles")
			c.JSON(200, gin.H{
				"username": username,
				"roles":    roles,
			})
		})
	}

	r.Run(":9999")
}
