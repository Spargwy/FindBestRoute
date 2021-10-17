package main

import (
	"strconv"
	"strings"
)

func ParseTimetableField(timetableString string) (timeTableField TimetableField, err error) {
	timetableString = strings.ReplaceAll(timetableString, " ", "")
	timetableFieldsValues := strings.Split(timetableString, ";")
	timeTableField.Cost, err = strconv.ParseFloat(timetableFieldsValues[3], 64)
	if err != nil {
		return
	}
	timeTableField.Train, timeTableField.FromStation,
		timeTableField.ToStation, timeTableField.FromTime,
		timeTableField.ToTime = timetableFieldsValues[0], timetableFieldsValues[1],
		timetableFieldsValues[2], timetableFieldsValues[4], timetableFieldsValues[5]
	timeTableField.ElapsedTime, err = ParseTimeToMinutes(timeTableField.FromTime, timeTableField.ToTime)
	if err != nil {
		return
	}
	return
}

// Calculating elapse time in minutes
func ParseTimeToMinutes(fromTime, toTime string) (elapsedTime int64, err error) {
	fromTimeValues := strings.Split(fromTime, ":")
	toTimeValues := strings.Split(toTime, ":")

	fromHour, err := strconv.ParseInt(fromTimeValues[0], 10, 32)
	if err != nil {
		return
	}
	fromMinute, err := strconv.ParseInt(fromTimeValues[1], 10, 32)
	if err != nil {
		return
	}

	toHour, err := strconv.ParseInt(toTimeValues[0], 10, 32)
	if err != nil {
		return
	}
	toMinute, err := strconv.ParseInt(toTimeValues[1], 10, 32)
	if err != nil {
		return
	}
	if toHour > fromHour {
		elapsedTime = (toHour - fromHour) * 60
	} else {
		elapsedTime = (toHour - fromHour + 24) * 60
	}
	elapsedTime += (toMinute - fromMinute)
	return
}

func ElementInArray(station string, path []TimetableField) bool {
	elementInArray := false
	for i := range path {
		if path[i].FromStation == station {
			elementInArray = true
			break
		}

	}
	return elementInArray
}
