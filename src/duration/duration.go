package duration

import (
	. "../models"
	"fmt"
	"time"
)

func ConvertYearMonthDayToTime(yearmonthday YearMonthDay) time.Time {
	return time.Date(yearmonthday.Year, time.Month(yearmonthday.Month), yearmonthday.Day, 0, 0, 0, 0, time.UTC)
}

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


func GetMinutesFromDuration(duration time.Duration) int {
	return int(duration.Minutes())
}

func GetHoursFromDuration(duration time.Duration) int {
	return int(duration.Hours())
}

func ConvertHoursToDays(hours int) int {
	days := hours/24
	return days
}

func ConvertDaysToWeeks(days int) WeekDay {
	weeks := days / 7
	daysInWeek := days % 7
	return WeekDay{
		Weeks: weeks,
		Days:  daysInWeek,
	}
}
