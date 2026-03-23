package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Activity int8

const (
	NoActivity Activity = iota
	Walking
	Jogging
	Cycling
)

type ActivityPost struct {
	Id              int64    `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	DurationHours   int8     `json:"durationHours"`
	DurationMinutes int8     `json:"durationMinutes"`
	DurationSeconds int8     `json:"durationSeconds"`
	Activity        Activity `json:"activity"`
}

var testData = []ActivityPost{
	{
		Id:              0,
		Title:           "title0",
		Description:     "description0",
		DurationHours:   0,
		DurationMinutes: 0,
		DurationSeconds: 0,
		Activity:        Walking,
	},
}

func main() {
	router := gin.Default()
	router.GET("/activities", getActivityPosts)
	router.GET("/activities/:id", getActivityPost)
	router.POST("/activities", postActivityPosts)
	router.Run("localhost:5000")
	fmt.Println("test")
}

func getActivityPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testData)
}

func getActivityPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, activity := range testData {
		if activity.Id == id {
			c.IndentedJSON(http.StatusOK, activity)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not found"})
}

func postActivityPosts(c *gin.Context) {
	var newActivity ActivityPost

	if err := c.BindJSON(&newActivity); err != nil {
		return
	}

	testData = append(testData, newActivity)
	c.IndentedJSON(http.StatusCreated, newActivity)
}
