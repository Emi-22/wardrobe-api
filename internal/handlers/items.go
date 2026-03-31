package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Emi-22/wardrobe-api/api/internal/db"
	"github.com/Emi-22/wardrobe-api/api/internal/models"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT id, name, classification, color, brand, favorite, created_at
		FROM closet_items
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Classification,
			&item.Color,
			&item.Brand,
			&item.Favorite,
			&item.CreatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var item models.Item
	err := db.DB.QueryRow(`
		SELECT *
		FROM closet_items
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.Name,
		&item.Classification,
		&item.Color,
		&item.Brand,
		&item.Favorite,
		&item.CreatedAt,
	)

	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow(`
		INSERT INTO closet_items (name, classification, color, brand, favorite)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`, item.Name, item.Classification, item.Color, item.Brand, item.Favorite).Scan(&item.ID, &item.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = db.DB.QueryRow(`
		SELECT id FROM closet_items WHERE id = $1
	`, id).Scan(new(int))
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	_, err = db.DB.Exec(`
		UPDATE closet_items
		SET name=$1, classification=$2, color=$3, brand=$4, favorite=$5
		WHERE id=$6
	`, item.Name, item.Classification, item.Color, item.Brand, item.Favorite, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.QueryRow(`
		SELECT id FROM closet_items WHERE id = $1
	`, params["id"]).Scan(new(int))
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(params["id"])

	_, err = db.DB.Exec(`
		DELETE FROM closet_items WHERE id = $1
	`, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusNoContent)
}
