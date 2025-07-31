package main

import (
	"encoding/json"
	"github.com/vorobyan/bookquest/model"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// Route for main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("BookQuest стартует сегодня! 🔥"))
		if err != nil {
			log.Fatal(err)
		}
	})

	// Route for achievements
	http.HandleFunc("/achievements", func(w http.ResponseWriter, r *http.Request) {

		// Create exemplar
		firstAchievement := model.Achievement{
			Name: "Git создан и привязан, спасибо Goland!",
			XP:   15,
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

	// Обработчик статики
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route for html page with achievements
	http.HandleFunc("/achievements/page", func(w http.ResponseWriter, r *http.Request) {

		//Parsing template
		tmpl, err := template.ParseFiles(filepath.Join("templates", "achievement.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := model.PageData{
			Title: "Твоя первая ачивка!",
			Achievement: model.Achievement{
				Name: "Git создан и привязан, спасибо Goland!",
				XP:   15,
				Icon: "🤙",
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
