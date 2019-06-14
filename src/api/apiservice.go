package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestDate struct {
	Start YearMonthDay 
	End YearMonthDay
}

type YearMonthDay struct {
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
}

type MonthDay struct {
	Months int `json:"months"`
	Days int `json:"days"`
}

type WeekDay struct {
	Weeks int `json:"weeks"`
	Days int `json:"days"`
}

type Duration struct {
	Days int `json:"days"`
	StartFullDate string `json:"startFullDate"`
	EndFullDate string `json:"endFullDate"`
	DaysToMonths MonthDay `json:"daysToMonths"`
	DaysToWeeks WeekDay `json:"daysToWeeks"`
	Seconds int `json:"seconds"`
	Minutes int `json:"minutes"`
	Hours int `json:"hours"`
	
	
}

func DurationBetweenTimes(writer http.ResponseWriter, request *http.Request)  {
		duration := Duration{
			StartFullDate : "Thursday, 16 October 1997",
			EndFullDate : "Monday, 10 June 2019",
			DaysToMonths : MonthDay{
				Months:259,
				Days:25,
			},
			DaysToWeeks: WeekDay{
				Weeks:1129,
				Days:4,
			},
			Days : 7907,
			Seconds:683164800,
			Minutes:11386080,
			Hours:189768,
		}
		durationjson,_ := json.Marshal(duration)
		fmt.Fprint(writer, string(durationjson))
}
