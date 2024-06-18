package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var i int

type counterResponse struct {
	Message string `json:"message"`
	Counter int    `json:"counter"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WEB server cyk")
	})

	http.HandleFunc("/senddata", recieveData)

	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Counter called")
		response := counterResponse{
			Message: "counter indicating how many times has this endpoint been called",
			Counter: i,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		i++
	})

	fmt.Printf("Starting server at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

func recieveData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var responseData []float32
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(responseData)
	err = saveData(responseData)
	if err != nil {
		http.Error(w, "Error saving data to db", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

}

func saveData(data []float32) error {
	path := "/Users/mateuszbialkowski/Desktop/programowanie/golang_sandbox/go-web-data-plotter/db/measurements.txt"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	stringSlice := make([]string, len(data))
	for i, v := range data {
		stringSlice[i] = strconv.FormatFloat(float64(v), 'f', -1, 32)
	}
	stringToSave := strings.Join(stringSlice, ", ")
	stringToSave = stringToSave + "\n"
	_, err = file.WriteString(stringToSave)
	if err != nil {
		fmt.Println("write error:", err)
		return err
	}
	file.Close()
	return nil
}
