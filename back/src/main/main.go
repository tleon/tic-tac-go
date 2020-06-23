package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var grid Grid

func getGrid(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewEncoder(w).Encode(grid.content)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/get-grid", getGrid)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}