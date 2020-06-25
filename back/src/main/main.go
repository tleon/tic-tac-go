package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var grid Grid

func getGrid(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(grid.content); err != nil {
		fmt.Println(err)
	}
}

func clicked(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	var data map[string]interface{}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)

	if err := json.Unmarshal([]byte(buf.String()), &data); err != nil {
		fmt.Println(err)
	}

	if data["col"] != nil && data["row"] != nil {
		grid.HandlePlay(int(data["row"].(float64)), int(data["col"].(float64)), 1)

		if err := json.NewEncoder(w).Encode(grid.content); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/get-grid", getGrid)
	http.HandleFunc("/clicked", clicked)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
