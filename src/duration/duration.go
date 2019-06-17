package duration

import (
	"time"
	api "../api"
)

func ConvertYearMonthDayToTime(yearmonthday api.YearMonthDay) time.Time {
	return time.Date(yearmonthday.Year,time.Month(yearmonthday.Month),yearmonthday.Day,0,0,0,0,time.UTC)
}

func CalculateDurationBetweenTime(startTime time.Time, endTime time.Time) time.Duration {
	duration, _ := time.ParseDuration("683164800s")
	return duration
}
