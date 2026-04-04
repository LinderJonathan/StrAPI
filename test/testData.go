package test

import "strAPI/util"

var TestData = []util.Activity{
	{
		Id:              0,
		Title:           "title0",
		Description:     "description0",
		DurationHours:   0,
		DurationMinutes: 0,
		DurationSeconds: 0,
		ActivityType:    util.ActivityType(util.Walking),
	},
	{
		Id:              1,
		Title:           "title1",
		Description:     "description1",
		DurationHours:   1,
		DurationMinutes: 1,
		DurationSeconds: 1,
		ActivityType:    util.ActivityType(util.Jogging),
	},
	{
		Id:              2,
		Title:           "title2",
		Description:     "description2",
		DurationHours:   2,
		DurationMinutes: 2,
		DurationSeconds: 2,
		ActivityType:    util.ActivityType(util.Cycling),
	},
}
