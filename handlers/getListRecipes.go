package handlers

import (
	"encoding/json"
	"net/http"
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
    query += " AND r.title LIKE ? COLLATE NOCASE"
    args = append(args, "%"+title+"%")
	}

	if cuisine != "" {
		query += " AND r.cuisine = ?"
		args = append(args, cuisine)
	}

	if ingredient != "" {
		query += " AND i.name LIKE ? COLLATE NOCASE"
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