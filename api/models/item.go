package models

import "time"

type Item struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Classification string    `json:"classification"`
	Color          string    `json:"color"`
	Brand          string    `json:"brand"`
	Favorite       bool      `json:"favorite"`
	CreatedAt      time.Time `json:"createdat"`
}
