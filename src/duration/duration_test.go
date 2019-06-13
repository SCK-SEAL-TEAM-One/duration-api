package duration

import (
	"testing"
)

func TestInputJsonText_Year_1997_Month_10_Day_16_Should_Year_1997_Month_10_Day_16(t *testing.T){
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

func TestInputJsonText_Year_1997_Month_10_Day_27_Should_Year_1997_Month_10_Day_27(t *testing.T){
	expectedResultYear := 1997
	expectedResultMonth := 10
	expectedResultDay := 27

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"year":1997,"month":10,"day":27}`)


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

func TestInputJsonText_Year_1996_Month_2_Day_5_Should_Year_1996_Month_2_Day_27(t *testing.T){
	expectedResultYear := 1996
	expectedResultMonth := 2
	expectedResultDay := 5

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"year":1996,"month":2,"day":5}`)


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

func TestInputJsonText_Year_2019_Month_6_Day_10_Should_Year_2019_Month_6_Day_5(t *testing.T){
	expectedResultYear := 2019
	expectedResultMonth := 6
	expectedResultDay := 10

	actualResultYear, actualResultMonth, actualResultDay := GetYearMonthDayFromRequestDate(`{"year":2019,"month":6,"day":10}`)


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