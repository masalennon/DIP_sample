package main

import (
	"encoding/json"
	"net/http"

	"github.com/masalennon/DIP_sample/db"
	"github.com/masalennon/DIP_sample/store"
)

func main() {
	db.Init()
	http.Handle("/employees/", http.StripPrefix("/employees/", http.HandlerFunc(getEmployee)))
	http.ListenAndServe(":8080", nil)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	employee, err := store.GetEmployeeByID(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(employee)
}
