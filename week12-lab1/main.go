package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// กำหนด Claims structure
type CustomClaims struct {
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

var secretKey = []byte("my-super-secret-key-change-in-production")

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
func main() {
	// สร้าง token
	token, err := generateToken(
		1,
		"alice",
		[]string{"admin"},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated Token >>")
	fmt.Println(token)
	fmt.Println()

	// ตรวจสอบ token
	claims, err := verifyToken(token)
	if err != nil {
		fmt.Println("Invalid token:", err)
		return
	}

	fmt.Println("Token verified successfully!")
	fmt.Printf("User ID >> %d\n", claims.UserID)
	fmt.Printf("Username >> %s\n", claims.Username)
	fmt.Printf("Roles >> %v\n", claims.Roles)
	fmt.Printf("Expires At >> %v\n", claims.ExpiresAt.Time)
	fmt.Printf("Issued At >> %v\n", claims.IssuedAt.Time)
	fmt.Printf("Issuer >> %v\n", claims.Issuer)

	// ทดสอบ token ที่ถูกแก้ไข
	tamperedToken := token[:len(token)-5] + "xxxxx"
	_, err = verifyToken(tamperedToken)
	if err != nil {
		fmt.Println("\nTampered token rejected", err)
	}
}
