package main

import (
	"encoding/json"
	"net/http"

	"github.com/masalennon/DIP_sample/employee"

	"github.com/masalennon/DIP_sample/db"
	"github.com/masalennon/DIP_sample/store"
)

func main() {
	d := db.Init()
	es := store.NewEmployeeGormStore(d)
	h := NewEmployeeHandler(es)
	http.Handle("/employees/", http.StripPrefix("/employees/", http.HandlerFunc(h.getEmployee)))
	http.ListenAndServe(":8080", nil)
}

type EmployeeHandler struct {
	es employee.Store
}

func NewEmployeeHandler(es employee.Store) *EmployeeHandler {
	return &EmployeeHandler{
		es: es,
	}
}

func (h EmployeeHandler) getEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	employee, err := h.es.GetEmployeeByID(id)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(employee)
}
