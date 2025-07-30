package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Route for main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BookQuest стартует сегодня! 🔥"))
	})

	// Route for achievements
	http.HandleFunc("/achievements", func(w http.ResponseWriter, r *http.Request) {
		// Create structure
		type Achievement struct {
			Name string `json:"name"`
			XP   int    `json:"xp"`
		}

		// Create exemplar
		firstAchievement := Achievement{
			Name: "Первая ачивка",
			XP:   10,
		}

		// Coding to JSON
		jsonData, _ := json.Marshal(firstAchievement)

		// Sending request
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(jsonData)
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
