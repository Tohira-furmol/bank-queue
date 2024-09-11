package controllers

import (
	"net/http"
	"time"

	"bank-queue-system/models"
	"bank-queue-system/services"

	"github.com/gin-gonic/gin"
)

var queues []models.Queue
var users []models.User

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func BookQueue(c *gin.Context) {
	var newQueue models.Queue
	if err := c.BindJSON(&newQueue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newQueue.ID = uint(len(queues) + 1)
	newQueue.TicketNo = generateTicketNo()
	newQueue.CreatedAt = time.Now()
	queues = append(queues, newQueue)

	c.JSON(http.StatusCreated, newQueue)
}

func generateTicketNo() string {
	return "Q" + time.Now().Format("150405")
}

func BookQueueController(c *gin.Context) {
	var newQueue models.Queue
	if err := c.ShouldBindJSON(&newQueue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdQueue, err := services.BookQueue(newQueue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdQueue)
}