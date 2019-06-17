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