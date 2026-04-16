package main

import (
	"fmt"
	"net/http"
	"strconv"

	"strAPI/db"
	"strAPI/test"
	"strAPI/util"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	defer db.DBConn.Close()

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

	var activities []util.Activity

	query := `
	SELECT 
		id, 
		title,
		description,
		durationHours,
		durationMinutes,
		durationSeconds,
		activityType
	FROM Activities 
	`
	rows, err := db.DBConn.Query(query)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var activity util.Activity
		err := rows.Scan(
			&activity.Id,
			&activity.Title,
			&activity.Description,
			&activity.DurationHours,
			&activity.DurationMinutes,
			&activity.DurationSeconds,
			&activity.ActivityType)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		activities = append(activities, activity)
	}

	c.IndentedJSON(http.StatusOK, activities)
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

	var activity util.Activity
	query := `
	SELECT 
		id, 
		title,
		description,
		durationHours,
		durationMinutes,
		durationSeconds,
		activityType
	FROM Activities 
	WHERE id = ?
	`

	err = db.DBConn.QueryRow(query, id).Scan(
		&activity.Id,
		&activity.Title,
		&activity.Description,
		&activity.DurationHours,
		&activity.DurationMinutes,
		&activity.DurationSeconds,
		&activity.ActivityType)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, activity)
}

/*
 * Function postActivity
 *
 * POST request.
 * Creates a new endpoint with a given Id.
 * Successful if the generated Id is unique and validation passes
 */
func postActivity(c *gin.Context) {

	// request struct. does not contain ID
	var activityRequest util.ActivityRequest

	if err := c.BindJSON(&activityRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := util.ValidateActivity(&activityRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
	INSERT INTO Activities
		(title, description, durationHours, durationMinutes, durationSeconds, activityType)
	VALUES(?, ?, ?, ?, ?, ?)
	`

	result, err := db.DBConn.Exec(
		query,
		&activityRequest.Title,
		&activityRequest.Description,
		&activityRequest.DurationHours,
		&activityRequest.DurationMinutes,
		&activityRequest.DurationSeconds,
		&activityRequest.ActivityType)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch id "})
	}

	// add id to response struct, Signal client of the new endpoint
	var activity util.Activity = util.Activity{
		Id:              uint64(id),
		Title:           activityRequest.Title,
		Description:     activityRequest.Description,
		DurationHours:   activityRequest.DurationHours,
		DurationMinutes: activityRequest.DurationMinutes,
		DurationSeconds: activityRequest.DurationSeconds,
		ActivityType:    activityRequest.ActivityType,
	}

	c.Header("Location", fmt.Sprintf("/activities/%d"))

	c.IndentedJSON(http.StatusCreated, activity)
}

/* Function putActivity
 *
 * PUT request
 * Matches the endpoint Id towards existing activities.
 * Successful if the Id exists, otherwise returns an error in the response
 */
func putActivity(c *gin.Context) {

	// TODO: putActivity shouldn't obtain an ID, just like the POST
	var newActivity util.Activity
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&newActivity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newActivity.Id = id // using ID from URL (not JSON body)

	// TODO: call validateActivity() on the new activity data
	query := `
	INSERT INTO activities
		(id, title, description, durationHours, durationMinutes, durationSeconds, activityType)
	VALUES(?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY
		title=VALUES(title)
		description=VALUES(description)
		durationHours=VALUES(durationHours)
		durationMinutes=VALUES(durationMinutes)
		durationSeconds=VALUES(durationSeconds)
		activityType=VALUES(activityType)
	`

	_, err = db.DBConn.Exec(
		query,
		&newActivity.Id,
		&newActivity.Title,
		&newActivity.Description,
		&newActivity.DurationHours,
		&newActivity.DurationMinutes,
		&newActivity.DurationSeconds,
		&newActivity.ActivityType)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Signal client of the new/modified endpoint
	c.Header("Location", fmt.Sprintf("/activities/%d", newActivity.Id))
	c.IndentedJSON(http.StatusOK, newActivity)
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
