package duration

import (
	"testing"
)

func TestInputJsonText_Year_1997_Month_10_Day_16_Should_Year_1997_Month_10_Year_Day_16(t *testing.T){
	expectedResultYear := 1997
	expectedResultMonth := 10
	expectedResultDay := 16

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"year":1997,"month":10,"day":16}`)


	if actualResultYear != expectedResultYear {
		t.Errorf("Expect(Year) %v but get %v",expectedResultYear,actualResultYear)
	}
	if actualResultMonth != expectedResultMonth {
		t.Errorf("Expect(Month) %v but get %v",expectedResultMonth,actualResultMonth)
	}
	if actualResultDay != expectedResultDay {
		t.Errorf("Expect(Day) %v but get %v",expectedResultDay,actualResultDay)
	}
}