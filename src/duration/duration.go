package duration

import (
	"encoding/json"
)

func GetYearMonthDayFromRequestDate(datejson string) (year int,month int,day int) {
	var jsonBlob = []byte(datejson)
	type Date struct{
		Day int `json:"day"`
		Month int `json:"month"`
		Year int `json:"year"`
		
		
	}
	var date Date
	_ = json.Unmarshal(jsonBlob,&date)
	return date.Year,date.Month,date.Day
}
