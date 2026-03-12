package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

var fakeItems = []Item{
	{
		ID:          "1",
		Name:        "Sample Item 1",
		Description: "This is a fake item for demonstration",
		CreatedAt:   time.Now().Add(-24 * time.Hour),
	},
	{
		ID:          "2",
		Name:        "Sample Item 2",
		Description: "Another fake item for testing",
		CreatedAt:   time.Now().Add(-48 * time.Hour),
	},
}

func main() {
	r := gin.Default()

	r.GET("/health", healthHandler)
	r.GET("/items", getItemsHandler)
	r.GET("/items/:id", getItemByIDHandler)
	r.POST("/items", createItemHandler)

	r.Run(":8080")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func getItemsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": fakeItems,
		"count": len(fakeItems),
	})
}

func getItemByIDHandler(c *gin.Context) {
	id := c.Param("id")

	for _, item := range fakeItems {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Item not found",
	})
}

func createItemHandler(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newItem := Item{
		ID:          string(rune(len(fakeItems) + 1)),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	fakeItems = append(fakeItems, newItem)

	c.JSON(http.StatusCreated, newItem)
}
