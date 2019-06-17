package duration

import (
	"testing"
)

func Test_GetYearMonthDayFromRequestDate_InputJsonText_Day_16_Month_10_Year_1997_Should_Year_1997_Month_10_Day_16(t *testing.T){
	expectedResultYear := 1997
	expectedResultMonth := 10
	expectedResultDay := 16

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"day":16,"month":10,"year":1997}`)


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

func Test_GetYearMonthDayFromRequestDate_InputJsonText_Day_27_Month_10_Year_1997_Should_Year_1997_Month_10_Day_27(t *testing.T){
	expectedResultYear := 1997
	expectedResultMonth := 10
	expectedResultDay := 27

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"day":27,"month":10,"year":1997}`)


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

func Test_GetYearMonthDayFromRequestDate_InputJsonText_Day_5_Month_2_Year_1996_Should_Year_1996_Month_2_Day_27(t *testing.T){
	expectedResultYear := 1996
	expectedResultMonth := 2
	expectedResultDay := 5

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"day":5,"month":2,"year":1996}`)


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

func Test_GetYearMonthDayFromRequestDate_InputJsonText_Day_10_Month_6_Year_2019_Should_Year_2019_Month_6_Day_5(t *testing.T){
	expectedResultYear := 2019
	expectedResultMonth := 6
	expectedResultDay := 10

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"day":10,"month":6,"year":2019}`)


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
