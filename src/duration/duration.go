package duration

import (
	api "../api"
	"fmt"
	"time"
)

func ConvertTimeToFullDate(time time.Time) string {
	weekDay := time.Weekday()
	year := time.Year()
	month := time.Month()
	day := time.Day()
	fullDate := fmt.Sprintf("%s, %v %v %v", weekDay, day, month, year)
	return fullDate
}

func ConvertYearMonthDayToTime(yearmonthday api.YearMonthDay) time.Time {
	return time.Date(yearmonthday.Year, time.Month(yearmonthday.Month), yearmonthday.Day, 0, 0, 0, 0, time.UTC)
}

func CalculateDurationBetweenTime(startTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}

func GetSecondsFromDuration(duration time.Duration) int {
	return int(duration.Seconds())
}


