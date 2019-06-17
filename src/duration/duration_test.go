package duration

import (
	api "../api"
	"testing"
	"time"
)

func Test_ConvertYearMonthDayToTime_1997_Year_10_Month_16_Day_Should_Time_Year_1997_Month_10_Day_16(t *testing.T) {
	yearmonthday := api.YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   16,
	}

	expectedResult := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)

	actualResult := ConvertYearMonthDayToTime(yearmonthday)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertYearMonthDayToTime_1997_Year_10_Month_27_Day_Should_Time_Year_1997_Month_10_Day_27(t *testing.T) {
	yearmonthday := api.YearMonthDay{
		Year:  1997,
		Month: 10,
		Day:   27,
	}

	expectedResult := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)

	actualResult := ConvertYearMonthDayToTime(yearmonthday)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertYearMonthDayToTime_1996_Year_2_Month_5_Day_Should_Time_Year_1996_Month_2_Day_5(t *testing.T) {
	yearmonthday := api.YearMonthDay{
		Year:  1996,
		Month: 2,
		Day:   5,
	}

	expectedResult := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)

	actualResult := ConvertYearMonthDayToTime(yearmonthday)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertYearMonthDayToTime_2019_Year_6_Month_10_Day_Should_Time_Year_2019_Month_6_Day_10(t *testing.T) {
	yearmonthday := api.YearMonthDay{
		Year:  2019,
		Month: 6,
		Day:   10,
	}

	expectedResult := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)

	actualResult := ConvertYearMonthDayToTime(yearmonthday)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertTimeToFullDate_Time_Day_16_Month_10_Year_1997_Should_Thursday_16_October_1997(t *testing.T) {
	time := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	expectedResult := "Thursday, 16 October 1997"
	actualResult := ConvertTimeToFullDate(time)
	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func TestCalculateDurationBetweenTime_Year_1997_Month_10_Day_16_And_Year_2019_Month_06_Day_10_Should_Get_Duration_683164800_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("683164800s")
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func TestCalculateDurationBetweenTime_Year_1997_Month_10_Day_27_And_Year_2019_Month_06_Day_10_Should_Get_Duration_682144000_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("682214400s")
	startTime := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func TestCalculateDurationBetweenTime_Year_1996_Month_02_Day_05_And_Year_2019_Month_06_Day_10_Should_Get_Duration_736646400_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("736646400s")
	startTime := time.Date(1996, 02, 05, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_683164800_Seconds_Should_683164800_Seconds(t *testing.T) {
	expectedResult := 683164800
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_682214400_Seconds_Should_682214400_Seconds(t *testing.T) {
	expectedResult := 682214400
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_736646400_Seconds_Should_736646400_Seconds(t *testing.T) {
	expectedResult := 736646400
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMinutesFromDuration_Duration_683164800_Seconds_Should_11386080_Minutes(t *testing.T) {
	expectedResult := 11386080
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMinutesFromDuration_Duration_682214400_Seconds_Should_11370240_Minutes(t *testing.T) {
	expectedResult := 11370240
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMinutesFromDuration_Duration_736646400_Seconds_Should_12277440_Minutes(t *testing.T) {
	expectedResult := 12277440
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetHoursFromDuration_Duration_683164800_Seconds_Should_189768_Hours(t *testing.T) {
	expectedResult := 189768
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetHoursFromDuration_Duration_682214400_Seconds_Should_189504_Hours(t *testing.T) {
	expectedResult := 189504
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetHoursFromDuration_Duration_736646400_Seconds_Should_204624_Hours(t *testing.T) {
	expectedResult := 204624
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}


