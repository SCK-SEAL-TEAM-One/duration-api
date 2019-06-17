package api

import (
	du "../duration"
	. "../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DurationBetweenTimes(writer http.ResponseWriter, request *http.Request) {
	requestDate := GetYearMonthDayFromRequestDate(request)
	startTime := du.ConvertYearMonthDayToTime(requestDate.Start)
	endTime := du.ConvertYearMonthDayToTime(requestDate.End)
	durationValue := du.CalculateDurationBetweenTime(startTime, endTime)
	totalHours := du.GetHoursFromDuration(durationValue)
	totalDay := du.ConvertHoursToDays(totalHours)
	duration := Duration{
		StartFullDate: du.ConvertTimeToFullDate(startTime),
		EndFullDate:   du.ConvertTimeToFullDate(endTime),
		DaysToMonths:  du.CalculateMonthsBetweenStartTimeAndEndTime(requestDate.Start, requestDate.End),
		DaysToWeeks:   du.ConvertDaysToWeeks(totalDay),
		Days:          du.ConvertHoursToDays(totalHours),
		Seconds:       du.GetSecondsFromDuration(durationValue),
		Minutes:       du.GetMinutesFromDuration(durationValue),
		Hours:         du.GetHoursFromDuration(durationValue),
	}
	durationjson, _ := json.Marshal(duration)
	fmt.Fprint(writer, string(durationjson))
}
func GetYearMonthDayFromRequestDate(request *http.Request) RequestDate {
	body, _ := ioutil.ReadAll(request.Body)
	var requestDate RequestDate
	_ = json.Unmarshal(body, &requestDate)
	return requestDate
}
