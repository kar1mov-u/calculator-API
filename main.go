package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Base struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "THis is working")
	})

	mux.HandleFunc("POST /add", addition)
	mux.HandleFunc("POST /sub", subtarct)
	mux.HandleFunc("POST /multiply", multiply)
	mux.HandleFunc("POST /divide", divide)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

// Hadnlers
func addition(w http.ResponseWriter, r *http.Request) {
	help(w, r, "add")
}
func subtarct(w http.ResponseWriter, r *http.Request) {
	help(w, r, "sub")
}
func multiply(w http.ResponseWriter, r *http.Request) {
	help(w, r, "mul")
}
func divide(w http.ResponseWriter, r *http.Request) {
	help(w, r, "div")
}

// Function that does most of the logic
func help(w http.ResponseWriter, r *http.Request, t string) {
	var data Base

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Json: ", http.StatusBadRequest)
		return
	}

	if data.Num1 == 0 && data.Num2 == 0 {
		http.Error(w, "Both values cannot be 0 or missing", http.StatusBadRequest)
		return
	}

	resp := map[string]float64{"result": 0}
	if t == "add" {
		resp["result"] = data.Num1 + data.Num2
	} else if t == "sub" {
		resp["result"] = data.Num1 - data.Num2
	} else if t == "mul" {
		resp["result"] = data.Num1 * data.Num2
	} else {
		if data.Num2 == 0 {
			http.Error(w, "Cannot divide by 0", http.StatusBadRequest)
			return
		}
		resp["result"] = data.Num1 / data.Num2
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Cannot encode JSON", http.StatusInternalServerError)
	}

}
