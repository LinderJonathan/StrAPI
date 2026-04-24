package db

import (
	"database/sql"
	"log"
	"strconv"

	"strAPI/util"

	_ "modernc.org/sqlite"
)

var DBConn *sql.DB

/*
 * Function InitDB
 *
 * Initializes the SQLite database.
 */
func InitDB() {
	var err error
	DBConn, err = sql.Open("sqlite", "./db/app.db")

	if err != nil {
		log.Fatalln(err)
	}

	DBConn.SetConnMaxLifetime(0)
	DBConn.SetConnMaxIdleTime(0)
	DBConn.SetMaxOpenConns(1)
	DBConn.SetMaxIdleConns(1)

	if DBConn == nil {
		log.Fatalln("DB is nil")
	}
	createActivityTable()
}

/*
 * Function createActivityTable
 *
 * Creates an SQL table for activities.
 * Only creates a table if does not already exist
 */
func createActivityTable() {
	var err error
	numUniqueActivitesStr := strconv.Itoa(int(util.NumberOfActivites))
	query := `
	CREATE TABLE IF NOT EXISTS Activities (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		durationHours INTEGER CHECK (durationHours >= 0),
		durationMinutes INTEGER CHECK (durationMinutes >= 0),
		durationSeconds INTEGER CHECK (durationSeconds >= 0),
		activityType INTEGER CHECK (activityType BETWEEN 0 AND ` + numUniqueActivitesStr + `)
	);
	`
	_, err = DBConn.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
