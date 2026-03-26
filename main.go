package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityType int8

const (
	NoActivity ActivityType = iota
	Walking
	Jogging
	Cycling
)

type Activity struct {
	Id              int64        `json:"id"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	DurationHours   int8         `json:"durationHours"`
	DurationMinutes int8         `json:"durationMinutes"`
	DurationSeconds int8         `json:"durationSeconds"`
	ActivityType    ActivityType `json:"activity"`
}

func main() {
	router := gin.Default()
	router.GET("/activities", getAllActivities)
	router.GET("/activities/:id", getActivity)
	router.POST("/activities", postActivity)
	router.PUT("/activities/:id", putActivity)
	router.Run("localhost:5000")
}

func getAllActivities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testData)
}

func getActivity(c *gin.Context) {
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

// TODO:
// Check that Id is unique
func postActivity(c *gin.Context) {
	var newActivity Activity

	if err := c.BindJSON(&newActivity); err != nil {
		return
	}

	testData = append(testData, newActivity)
	c.IndentedJSON(http.StatusCreated, newActivity)
}

// put towards specific id.
// only works if the endpoint we put towards has data already
func putActivity(c *gin.Context) {

	var newActivity Activity
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, activity := range testData {
		if activity.Id == id {
			// Id exists, bind data to new variable and replace new values?
			if err := c.BindJSON(&newActivity); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			testData[i] = newActivity
			c.IndentedJSON(http.StatusOK, newActivity)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}
}

func deleteActivity(c *gin.Context)
