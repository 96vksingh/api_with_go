package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip"`
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/all", getalldata)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hellllo")

}

func getalldata(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
		{Name: "vishal", City: "gkp", Zipcode: "230983"},
		{Name: "vishal2", City: "gkp2", Zipcode: "230984"},
		{Name: "vishal3", City: "gkp3", Zipcode: "230985"},
		{Name: "vishal4", City: "gkp4", Zipcode: "230986"},
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
