package util

import "errors"

type ActivityType int8

const (
	NoActivity ActivityType = iota
	Walking
	Jogging
	Cycling
	NumberOfActivites // Used to retrieve amount of activites
)

type Activity struct {
	Id              uint64       `json:"id"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	DurationHours   uint8        `json:"durationHours"`
	DurationMinutes uint8        `json:"durationMinutes"`
	DurationSeconds uint8        `json:"durationSeconds"`
	ActivityType    ActivityType `json:"ActivityType"`
}

type ActivityRequest struct {
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	DurationHours   uint8        `json:"durationHours"`
	DurationMinutes uint8        `json:"durationMinutes"`
	DurationSeconds uint8        `json:"durationSeconds"`
	ActivityType    ActivityType `json:"activityType"`
}

func ValidateActivity(activityRequest *ActivityRequest) error {

	if activityRequest.Title == "" {
		return errors.New("Activities are required to have a title")
	} else if activityRequest.ActivityType == NoActivity {
		return errors.New("Activities are required to be of a certain type")
	}

	return nil
}
