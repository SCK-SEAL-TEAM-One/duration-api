package duration

import (
	"testing"
	"time"
)

func TestCalculateDurationBetweenTime_Year_1997_Month_10_Day_16_And_Year_2019_Month_06_Day_10_Should_Get_Duration_683164800_Secounds(t *testing.T) {
	expectedResult, _ := time.ParseDuration("683164800s")
	startTime := time.Date(1997, 10, 16, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2019, 06, 10, 0, 0, 0, 0, time.UTC)

	actualResult := CalculateDurationBetweenTime(startTime, endTime)

	if expectedResult != actualResult {
		t.Errorf("need %v but got %v", expectedResult, actualResult)
	}
}
