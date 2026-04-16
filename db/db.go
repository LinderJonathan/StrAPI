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
		id INT AUTO_INCREMENT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		durationHours INT CHECK (durationHours >= 0),
		durationMinutes INT CHECK (durationMinutes >= 0),
		durationSeconds INT CHECK (durationSeconds >= 0),
		activityType INT CHECK (activityType BETWEEN 0 AND ` + numUniqueActivitesStr + `)
	);
	`
	_, err = DBConn.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
