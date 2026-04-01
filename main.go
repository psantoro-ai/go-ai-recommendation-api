package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}

var items = []Item{
	{"AI for Marketing", "marketing"},
	{"Python for Beginners", "programming"},
	{"Advanced AI Systems", "ai"},
	{"Digital Business Strategy", "business"},
	{"Machine Learning Basics", "ai"},
}

func recommendHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var results []Item

	for _, item := range items {
		if item.Category == category {
			results = append(results, item)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/recommend", recommendHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
