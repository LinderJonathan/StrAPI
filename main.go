package main

import (
	"fmt"
	"net/http"
	"strconv"

	"strAPI/test"
	"strAPI/util"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.GET("/activities", getAllActivities)
	router.GET("/activities/:id", getActivity)
	router.POST("/activities", postActivity)
	router.PUT("/activities/:id", putActivity)
	router.DELETE("/activities/:id", deleteActivity)
	router.Run("localhost:5000")
}

/*
 * Function getAllActivities
 *
 * GET request.
 * Returns all activities. This functionality is only for testing purposes
 */
func getAllActivities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, test.TestData)
}

/*
 * Function getActivity
 *
 * GET request.
 * Returns the respective activity based on the Id provided in the request
 * Successful if an activity with that Id exists, otherwise returns an
 * error in the response
 */
func getActivity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, activity := range test.TestData {
		if activity.Id == id {
			c.IndentedJSON(http.StatusOK, activity)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not found"})
}

/*
 * Function postActivity
 *
 * POST request.
 * Creates a new endpoint with a given Id.
 * Successful if the generated Id is unique and validation passes
 */
func postActivity(c *gin.Context) {

	// TODO: After creating a Id, check if its a duplicate

	var newActivity util.Activity

	if err := c.BindJSON(&newActivity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := util.ValidateActivity(&newActivity)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	test.TestData = append(test.TestData, newActivity)
	c.Header("Location", fmt.Sprintf("/activities/%d", newActivity.Id))
	// send response back to client
	c.IndentedJSON(http.StatusCreated, newActivity)
}

/* Function putActivity
 *
 * PUT request
 * Matches the endpoint Id towards existing activities.
 * Successful if the Id exists, otherwise returns an error in the response
 */
func putActivity(c *gin.Context) {

	var newActivity util.Activity
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: call validateActivity() on the new activity data

	for i, activity := range test.TestData {
		if activity.Id == id {
			// Id exists, bind data to new variable and replace activity data
			if err := c.BindJSON(&newActivity); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			test.TestData[i] = newActivity
			c.IndentedJSON(http.StatusOK, newActivity)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not found"})
	return
}

/* Function deleteActivity
 *
 * DELETE request.
 * Matches the endpoint Id towards existing activities.
 * Successful if the Id exists, otherwise returns an error the response
 */
func deleteActivity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, activity := range test.TestData {
		if activity.Id == id {
			test.TestData = append(test.TestData[:i], test.TestData[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"success": "deleted activity", "Id": id})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not found"})
}
