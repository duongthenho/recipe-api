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

func GetRecipes(w http.ResponseWriter, r *http.Request) {

	title := strings.TrimSpace(r.URL.Query().Get("title"))
	cuisine := strings.TrimSpace(r.URL.Query().Get("cuisine"))
	ingredient := strings.TrimSpace(r.URL.Query().Get("ingredient"))

	query := `
		SELECT DISTINCT r.id, r.title, r.image_url, r.cuisine, r.views
		FROM recipes r
		LEFT JOIN ingredients i ON r.id = i.recipe_id
		WHERE 1=1
	`
	args := []any{}

	if title != "" {
		query += " AND r.title LIKE ?"
		args = append(args, "%"+title+"%")
	}

	if cuisine != "" {
		query += " AND r.cuisine LIKE ?"
		args = append(args, "%"+cuisine+"%")
	}

	if ingredient != "" {
		query += " AND i.name LIKE ?"
		args = append(args, "%"+ingredient+"%")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var result []models.Recipe

	for rows.Next() {

		var rcp models.Recipe

		err := rows.Scan(
			&rcp.ID,
			&rcp.Title,
			&rcp.Image,
			&rcp.Cuisine,
			&rcp.Views,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// ingredients
		ingRows, _ := db.DB.Query(
			`SELECT name FROM ingredients WHERE recipe_id = ?`,
			rcp.ID,
		)
		for ingRows.Next() {
			var name string
			ingRows.Scan(&name)
			rcp.Ingredients = append(rcp.Ingredients, name)
		}
		ingRows.Close()

		result = append(result, rcp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

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

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// /deleteById/12
	idStr := strings.TrimPrefix(r.URL.Path, "/deleteById/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// xóa bảng con trước
	_, err = tx.Exec(`DELETE FROM ingredients WHERE recipe_id = ?`, id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = tx.Exec(`DELETE FROM steps WHERE recipe_id = ?`, id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := tx.Exec(`DELETE FROM recipes WHERE id = ?`, id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 500)
		return
	}

	aff, _ := res.RowsAffected()
	if aff == 0 {
		tx.Rollback()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"id":      id,
	})
}