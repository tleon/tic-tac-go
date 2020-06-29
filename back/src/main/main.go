package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var grid Grid

func getGrid(w http.ResponseWriter, req *http.Request,  _ httprouter.Params) {
	if err := json.NewEncoder(w).Encode(grid.content); err != nil {
		fmt.Println(err)
	}
}

func clicked(w http.ResponseWriter, req *http.Request,  _ httprouter.Params) {
	var coords Coords

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&coords)

	if err != nil {
		fmt.Println(err)
	}

	grid.HandlePlay(coords.Row, coords.Col, 1)
	grid.cellChecked++
	if grid.cellChecked <= 8 {
		makeIAPlay()
	}

	if err := json.NewEncoder(w).Encode(grid.content); err != nil {
		fmt.Println(err)
	}
}

func makeIAPlay() {
	for {
		col := rand.Intn(3)
		row := rand.Intn(3)
		if false == grid.content[row][col].IsChecked {
			grid.content[row][col].IsChecked = true
			grid.content[row][col].ByPlayer = 2
			grid.cellChecked++
			break
		}
	}
}

func main() {
	router := httprouter.New()

	router.GET("/get-grid", getGrid)
	router.POST("/clicked", clicked)

	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		fmt.Println(err)
	}
}
