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

func TestCalculateDurationBetweenTime_Year_1997_Month_10_Day_16_And_Year_2019_Month_06_Day_10_Should_Get_Duration_683164800_Secounds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("683164800s")
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}
