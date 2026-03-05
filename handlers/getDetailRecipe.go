package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"recipe-api/db"
	"recipe-api/models"
)

func GetRecipeDetail(w http.ResponseWriter, r *http.Request) {

	// /recipes/1
	idStr := strings.TrimPrefix(r.URL.Path, "/recipes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	var recipe models.Recipe

	err = db.DB.QueryRow(`
		SELECT id, title, cuisine, image_url, views
		FROM recipes
		WHERE id = ?
	`, id).Scan(
		&recipe.ID,
		&recipe.Title,
		&recipe.Cuisine,
		&recipe.Image,
		&recipe.Views,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "not found", 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// ingredients
	iRows, _ := db.DB.Query(`
		SELECT name FROM ingredients WHERE recipe_id = ?
	`, recipe.ID)
	for iRows.Next() {
		var name string
		iRows.Scan(&name)
		recipe.Ingredients = append(recipe.Ingredients, name)
	}
	iRows.Close()

	// steps
	sRows, _ := db.DB.Query(`
		SELECT text, image FROM steps
		WHERE recipe_id = ?
		ORDER BY id
	`, recipe.ID)

	for sRows.Next() {
		var s models.Step
		sRows.Scan(&s.Text, &s.Image)
		recipe.Steps = append(recipe.Steps, s)
	}
	sRows.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}