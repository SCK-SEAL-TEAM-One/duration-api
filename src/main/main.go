package main

import (
	"log"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/duration", func (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type","application/json")
		fmt.Fprintf(w,
		`{
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
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}