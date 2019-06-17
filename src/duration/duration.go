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

func ConvertDaysToWeeks(days int) WeekDay {
	weeks := days / 7
	daysInWeek := days % 7
	return WeekDay{
		Weeks: weeks,
		Days:  daysInWeek,
	}
}

func CalculateMonthsBetweenStartTimeAndEndTime(start YearMonthDay,end YearMonthDay) MonthDay {
	diffYear :=  end.Year-start.Year
	diffMonth := diffYear * 12
	diffMonth = diffMonth + (end.Month - start.Month)
	diffDay := end.Day-start.Day

	if diffDay < 0 {
		diffDay += 31
		diffMonth--
	}

	return MonthDay{
		Months:diffMonth,
		Days:diffDay,
	}
}
