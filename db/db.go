package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/TheGolurk/infraApi/models"
	"github.com/TheGolurk/infraApi/utils"
	_ "github.com/lib/pq" // Postgres Drive
)

var (
	dbname = os.Getenv("db")
	pwd    = os.Getenv("pswd")
	host   = os.Getenv("host")
	user   = os.Getenv("user")
	port   = 5432
	url    = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=disable", host, port, dbname, user, pwd)
)

func GetDatabase(w http.ResponseWriter) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
		return nil
	}

	if err = db.Ping(); err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
		return nil
	}

	return db
}
