package duration

import (
	api "../api"
	"fmt"
	"time"
)

func ConvertYearMonthDayToTime(yearmonthday api.YearMonthDay) time.Time {
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

func ConvertDaysToWeeks(days int) api.WeekDay{
	weeks := days/7
	daysInWeek := days%7
	return api.WeekDay{
		Weeks:weeks,
		Days:daysInWeek,
	}
}

func CalculateMonthsBetweenStartTimeAndEndTime(start api.YearMonthDay,end api.YearMonthDay) api.MonthDay {
	diffYear :=  end.Year-start.Year
	diffMonth := diffYear * 12
	diffMonth = diffMonth + (end.Month - start.Month)
	diffDay := end.Day-start.Day

	if diffDay < 0 {
		diffDay += 31
		diffMonth--
	}

	return api.MonthDay{
		Months:diffMonth,
		Days:diffDay,
	}
}
