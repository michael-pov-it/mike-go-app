package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Define the data as a slice of structs
	data := []struct {
		ID        string `json:"id"`
		First     string `json:"first"`
		Last      string `json:"last"`
		Department string `json:"department"`
	}{
		{"101", "Susan", "Matthew", "HR"},
		{"102", "Bill", "Gates", "Finance"},
		{"103", "Prateek", "Singh", "Engineering"},
		{"104", "Rakesh", "Singh", "IT"},
	}

	// Define the handler function
	http.HandleFunc("/api/v1/employees", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Set CORS headers to allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		// Encode the data as JSON and write it as the response
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the server on port 8888
	log.Fatal(http.ListenAndServe(":8888", nil))
}
