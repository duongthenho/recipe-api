package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"recipe-api/db"
)

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// xóa bảng con trước
	_, err = tx.Exec(`DELETE FROM ingredients`)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = tx.Exec(`DELETE FROM steps`)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := tx.Exec(`DELETE FROM recipes`)
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
	})
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