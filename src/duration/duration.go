package duration

import (
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

func CalculateDurationBetweenTime(startTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}

func GetSecondsFromDuration(duration time.Duration) int {
	return int(duration.Seconds())
}
