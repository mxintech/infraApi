package api

import (
	"database/sql"
	"net/http"

	"github.com/TheGolurk/infraApi/db"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetDatabase()

	err := db.QueryRow("INSERT INTO")

	switch err {
	case sql.ErrNoRows:
	case nil:
	default:
	}
}
