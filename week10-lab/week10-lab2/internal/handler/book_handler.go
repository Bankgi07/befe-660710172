package handler

import "github.com/gin-gonic/gin"

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"title"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags Books
// @Produce  json
// @Param   id   path      int     true  "Book ID"
// @Success 200  {object}  Book
// @Failure 404  {object}  ErrorResponse
// @Router  /Book/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "title": "mastering golang"})
}
