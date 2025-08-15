package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	// 单向流，只能读一次 body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	body, err = io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Println(json.Marshal(values))
}

func main() {
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/params", queryParams)
	http.ListenAndServe(":8080", nil)
}
