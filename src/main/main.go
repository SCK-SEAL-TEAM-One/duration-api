package main

import (
	"../api"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/duration", api.DurationBetweenTimes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}