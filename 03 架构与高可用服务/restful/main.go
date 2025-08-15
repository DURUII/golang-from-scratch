package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 35}
	employeeDB["Rose"] = &Employee{"e-2", "Rose", 45}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if _, err := fmt.Fprint(w, "Welcome!\n"); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)
	if info, ok = employeeDB[qName]; !ok {
		_, _ = w.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		if _, err := fmt.Fprintf(w, "{\"error\":\"%s\"}", err); err != nil {
			log.Printf("Error writing error response: %v", err)
		}
		return
	}
	_, _ = w.Write(infoJson)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	// http://127.0.0.1:8082/employee/Mike
	router.GET("/employee/:name", GetEmployeeByName)
	log.Fatal(http.ListenAndServe(":8082", router))
}
