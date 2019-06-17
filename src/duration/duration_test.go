package duration

import (
	"time"
	"testing"
)

func Test_ConvertTimeToFullDate_Time_Day_16_Month_10_Year_1997_Should_Thursday_16_October_1997(t *testing.T)  {
	time := time.Date(1997,10,16,0,0,0,0,time.UTC)
	expectedResult := "Thursday, 16 October 1997"

	actualResult := ConvertTimeToFullDate(time)

	if actualResult != expectedResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_683164800_Seconds_Should_683164800_Seconds(t *testing.T){
	expectedResult := 683164800
	duration,_ := time.ParseDuration("683164800s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_682214400_Seconds_Should_682214400_Seconds(t *testing.T){
	expectedResult := 682214400
	duration,_ := time.ParseDuration("682214400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}

func Test_GetSecondsFromDuration_Duration_736646400_Seconds_Should_736646400_Seconds(t *testing.T){
	expectedResult := 736646400
	duration,_ := time.ParseDuration("736646400s")

	actualResult := GetSecondsFromDuration(duration)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}