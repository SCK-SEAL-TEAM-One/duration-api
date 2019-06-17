package duration

import (
	"time"
	api "../api"
)

func ConvertYearMonthDayToTime(yearmonthday api.YearMonthDay) time.Time {
	return time.Date(yearmonthday.Year,time.Month(yearmonthday.Month),yearmonthday.Day,0,0,0,0,time.UTC)
}

