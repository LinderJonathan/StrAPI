package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"strAPI/db"
	"strAPI/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	defer db.DBConn.Close()

	router := gin.Default()
	router.Use(cors.Default())
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

	var activityRequests []util.ActivityRequest

	query := `
	SELECT 
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
		var activityRequest util.ActivityRequest
		err := rows.Scan(
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
		activityRequests = append(activityRequests, activityRequest)
	}

	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, activityRequests)
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
		activityRequest.Title,
		activityRequest.Description,
		activityRequest.DurationHours,
		activityRequest.DurationMinutes,
		activityRequest.DurationSeconds,
		activityRequest.ActivityType)

	if err != nil {
		log.Fatalln(err)
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatalln(err)
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

	c.Header("Location", fmt.Sprintf("/activities/%d", id))

	c.IndentedJSON(http.StatusCreated, activity)
}

/* Function putActivity
 *
 * PUT request
 * Matches the endpoint Id towards existing activities.
 * Successful if the Id exists, otherwise returns an error in the response
 */
func putActivity(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
		return
	}

	var activity util.Activity

	if err := c.BindJSON(&activity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var activityRequest util.ActivityRequest = util.ActivityRequest{
		Title:           activity.Title,
		Description:     activity.Description,
		DurationHours:   activity.DurationHours,
		DurationMinutes: activity.DurationMinutes,
		DurationSeconds: activity.DurationSeconds,
		ActivityType:    activity.ActivityType,
	}

	if err := util.ValidateActivity(&activityRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result sql.Result

	updateQuery := `
	UPDATE Activities SET
		title = ?,
		description = ?,
		durationHours = ?,
		durationMinutes = ?,
		durationSeconds = ?,
		activityType = ?
	WHERE id = ?
	`

	result, err = db.DBConn.Exec(
		updateQuery,
		activity.Title,
		activity.Description,
		activity.DurationHours,
		activity.DurationMinutes,
		activity.DurationSeconds,
		activity.ActivityType,
		id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Fatalln(err)
		return
	}

	if rows == 0 {
		insertQuery := `
		INSERT INTO Activities
			(title, description, durationHours, durationMinutes, durationSeconds, activityType)
		VALUES (?, ?, ?, ?, ?, ?)
		`

		result, err = db.DBConn.Exec(
			insertQuery,
			activity.Title,
			activity.Description,
			activity.DurationHours,
			activity.DurationMinutes,
			activity.DurationSeconds,
			activity.ActivityType)

		id, err = result.LastInsertId()
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	activity.Id = uint64(id)

	// Signal client of the new/modified endpoint
	c.Header("Location", fmt.Sprintf("/activities/%d", id))
	c.IndentedJSON(http.StatusCreated, activity)
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

	query := "DELETE FROM Activities WHERE id = ?"

	result, err := db.DBConn.Exec(
		query,
		id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsCount, err := result.RowsAffected()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "could not determine the affected rows"})
	}

	if rowsCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "activity not removed (not found)"})
	}

	c.Status(http.StatusNoContent)
}
