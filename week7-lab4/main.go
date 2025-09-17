package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB() {
	var err error

	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s db_name=%s sslmode=disable", host, port, user, password, name)
	// fmt.Println(conSt)
	db, err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatalf("failed to open database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	log.Println("successfully connected to database")
}

func main() {
	initDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()

		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealty", "error": err})
			return
		}
		c.JSON(200, gin.H{"message": "healthy"})

	})
}
