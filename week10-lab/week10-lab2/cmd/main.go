package main

import (
	_ "week10-lab2/docs" // ให้ Swag สร้างเอกสารใน Folder docs โดยอัตโนมัติ

	"week10-lab2/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Simple API Example
// @version         1.0
// @description     This is a simple example of using Gin with Swagger.
// @host            localhost:9999
// @BasePath        /api/v1
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// Swagger endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User API routes
	api := r.Group("/api/v1")
	{
		api.GET("/Book/:id", handler.GetBookByID) // ใช้ Handler จากไฟล์ user_handler.go
	}

	// Start server
	r.Run(":9999")
}
