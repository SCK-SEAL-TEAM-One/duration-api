package duration

import (
	. "../models"
	"testing"
	"time"
)

// ========== ConvertYearMonthDayToTime ==========
func Test_ConvertYearMonthDayToTime_1997_Year_10_Month_16_Day_Should_Get_Time_Is_Year_1997_Month_10_Day_16(t *testing.T) {
	yearmonthday := YearMonthDay{
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

func Test_ConvertYearMonthDayToTime_1997_Year_10_Month_27_Day_Should_Get_Time_Is_Year_1997_Month_10_Day_27(t *testing.T) {
	yearmonthday := YearMonthDay{
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

func Test_ConvertYearMonthDayToTime_1996_Year_2_Month_5_Day_Should_Get_Time_Is_Year_1996_Month_2_Day_5(t *testing.T) {
	yearmonthday := YearMonthDay{
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

func Test_ConvertYearMonthDayToTime_2019_Year_6_Month_10_Day_Should_Get_Time_Is_Year_2019_Month_6_Day_10(t *testing.T) {
	yearmonthday := YearMonthDay{
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

// ========== ConvertTimeToFullDate ==========
func Test_ConvertTimeToFullDate_By_Time_Is_Year_1997_Month_10_Day_16_Should_Be_Thursday_16_October_1997(t *testing.T) {
	time := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	expectedResult := "Thursday, 16 October 1997"

	actualResult := ConvertTimeToFullDate(time)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_ConvertTimeToFullDate_By_Time_Is_Year_1997_Month_10_Day_27_Should_Be_Monday_27_October_1997(t *testing.T) {
	time := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 27 October 1997"

	actualResult := ConvertTimeToFullDate(time)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_ConvertTimeToFullDate_By_Time_Is_Year_1996_Month_2_Day_5_Should_Be_Monday_5_February_1996(t *testing.T) {
	time := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 5 February 1996"

	actualResult := ConvertTimeToFullDate(time)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_ConvertTimeToFullDate_By_Time_Is_Year_2019_Month_6_Day_10_Should_Be_Monday_10_June_2019(t *testing.T) {
	time := time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC)
	expectedResult := "Monday, 10 June 2019"

	actualResult := ConvertTimeToFullDate(time)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

// ========== CalculateDurationBetweenTime ==========
func Test_CalculateDurationBetweenTime_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_06_EndDay_10_Should_Get_Duration_683164800_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("683164800s")
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationBetweenTime_StartYear_1997_StartMonth_10_StartDay_27_And_EndYear_2019_EndMonth_06_EndDay_10_Should_Get_Duration_682144000_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("682214400s")
	startTime := time.Date(1997, 10, 27, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_CalculateDurationBetweenTime_StartYear_1996_StartMonth_02_StartDay_05_And_EndYear_2019_EndMonth_06_EndDay_10_Should_Get_Duration_736646400_Seconds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("736646400s")
	startTime := time.Date(1996, 02, 05, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

// ========== GetSecondsFromDuration ==========
func Test_GetSecondsFromDuration_By_Duration_Is_683164800_Seconds_Should_Get_683164800_Seconds(t *testing.T) {
	expectedResult := 683164800
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_682214400_Seconds_Should_Get_682214400_Seconds(t *testing.T) {
	expectedResult := 682214400
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_736646400_Seconds_Should_Get_736646400_Seconds(t *testing.T) {
	expectedResult := 736646400
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

// ========== GetMinutesFromDuration ==========
func Test_GetMinutesFromDuration_Duration_683164800_Seconds_Should_Get_11386080_Minutes(t *testing.T) {
	expectedResult := 11386080
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMinutesFromDuration_Duration_682214400_Seconds_Should_Get_11370240_Minutes(t *testing.T) {
	expectedResult := 11370240
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetMinutesFromDuration_Duration_736646400_Seconds_Should_Get_12277440_Minutes(t *testing.T) {
	expectedResult := 12277440
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetMinutesFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

// ========== GetHoursFromDuration ==========
func Test_GetHoursFromDuration_Duration_683164800_Seconds_Should_Get_189768_Hours(t *testing.T) {
	expectedResult := 189768
	duration, _ := time.ParseDuration("683164800s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetHoursFromDuration_Duration_682214400_Seconds_Should_Get_189504_Hours(t *testing.T) {
	expectedResult := 189504
	duration, _ := time.ParseDuration("682214400s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetHoursFromDuration_Duration_736646400_Seconds_Should_Get_204624_Hours(t *testing.T) {
	expectedResult := 204624
	duration, _ := time.ParseDuration("736646400s")

	actualResult := GetHoursFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

// ========== ConvertHoursToDays ==========
func Test_ConvertHoursToDays_189768_Hours_Should_Be_7907_Days(t *testing.T) {
	expectedResult := 7907

	actualResult := ConvertHoursToDays(189768)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertHoursToDays_189504_Hours_Should_Be_7896_Days(t *testing.T) {
	expectedResult := 7896

	actualResult := ConvertHoursToDays(189504)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertHoursToDays_204624_Hours_Should_Be_8526_Days(t *testing.T) {
	expectedResult := 8526

	actualResult := ConvertHoursToDays(204624)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but got %v", expectedResult, actualResult)
	}
}

// ========== ConvertDaysToWeeks ==========
func Test_ConvertDaysToWeeks_7907_Days_Should_Be_1129_Weeks_And_4_DaysInWeek(t *testing.T) {
	weekday := WeekDay{
		Weeks: 1129,
		Days:  4,
	}
	expectedResult := weekday

	actualResult := ConvertDaysToWeeks(7907)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertDaysToWeeks_7896_Days_Should_Be_1128_Weeks_And_0_DaysInWeek(t *testing.T) {
	weekday := WeekDay{
		Weeks: 1128,
		Days:  0,
	}
	expectedResult := weekday

	actualResult := ConvertDaysToWeeks(7896)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_ConvertDaysToWeeks_8526_Days_Should_1128_Weeks_And_0_DaysInWeek(t *testing.T) {
	weekday := WeekDay{
		Weeks: 1218,
		Days:  0,
	}
	expectedResult := weekday

	actualResult := ConvertDaysToWeeks(8526)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

// ========== CalculateMonthsBetweenStartTimeAndEndTime ==========
func Test_CalculateMonthsBetweenStartTimeAndEndTime_StartYear_1997_StartMonth_10_StartDay_16_And_EndYear_2019_EndMonth_06_EndDay_10_Should_259_Months_25_Days(t *testing.T) {
	starttime := YearMonthDay{
		Year:1997,
		Month:10,
		Day:16,
	}
	endtime := YearMonthDay{
		Year:2019,
		Month:6,
		Day:10,
	}
	expectedResult := MonthDay{
		Months:259,
		Days:25,
	}

	actualResult := CalculateMonthsBetweenStartTimeAndEndTime(starttime,endtime)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CalculateMonthsBetweenStartTimeAndEndTime_StartYear_1997_StartMonth_10_StartDay_27_And_EndYear_2019_EndMonth_06_EndDay_10_Should_259_Months_14_Days(t *testing.T) {
	starttime := YearMonthDay{
		Year:1997,
		Month:10,
		Day:27,
	}
	endtime := YearMonthDay{
		Year:2019,
		Month:6,
		Day:10,
	}
	expectedResult := MonthDay{
		Months:259,
		Days:14,
	}

	actualResult := CalculateMonthsBetweenStartTimeAndEndTime(starttime,endtime)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_CalculateMonthsBetweenStartTimeAndEndTime_StartYear_1996_StartMonth_2_StartDay_5_And_EndYear_2019_EndMonth_06_EndDay_10_Should_259_Months_14_Days(t *testing.T) {
	starttime := YearMonthDay{
		Year:1996,
		Month:2,
		Day:5,
	}
	endtime := YearMonthDay{
		Year:2019,
		Month:6,
		Day:10,
	}
	expectedResult := MonthDay{
		Months:280,
		Days:5,
	}

	actualResult := CalculateMonthsBetweenStartTimeAndEndTime(starttime,endtime)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
