package main

import (
	"log"
	"net/http"
	"os"

	"recipe-api/db"
	"recipe-api/handlers"
)

func main() {

	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	initSQL, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DB.Exec(string(initSQL))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/recipes", handlers.GetRecipes)
	http.HandleFunc("/recipes/", handlers.GetRecipeDetail)
	http.HandleFunc("/seed", handlers.SeedData)
	http.HandleFunc("/deleteById/", handlers.DeleteRecipe)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux)))
}
func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}