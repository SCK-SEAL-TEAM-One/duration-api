package duration

import (
	"time"
	"testing"
	api "../api"
)

func Test_ConvertYearMonthDayToTime_1997_Year_10_Month_16_Day_Should_Time_Year_1997_Month_10_Day_16(t *testing.T){
	yearmonthday := api.YearMonthDay{
		Year:1997,
		Month:10,
		Day:16,
	}

	expectedResult := time.Date(1997,10,16,0,0,0,0,time.UTC)

	actualResult := ConvertYearMonthDayToTime(yearmonthday)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v",expectedResult,actualResult)
	}
}