package api

import (
	"fmt"
	"net/http"
)

func DurationBetweenTimes(writer http.ResponseWriter, request *http.Request)  {
	_, _ = fmt.Fprint(writer, `{
		"days":7907,
		"startFullDate":"Thursday, 16 October 1997",
		"endFullDate":"Monday, 10 June 2019",
		"daysToMonths":{
			"months":259,
			"days":25
		},
		"daysToWeeks":{
			"weeks":1129,
			"days":4
		},
		"seconds":683164800,
		"minutes":11386080,
		"hours":189768
		}`)
}
