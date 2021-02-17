package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TheGolurk/infraApi/db"
	"github.com/TheGolurk/infraApi/models"
	"github.com/TheGolurk/infraApi/utils"

	_ "github.com/lib/pq" // Postgres Driver
)

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	db, err := db.GetDatabase()
	if err != nil {
		return errors.New(fmt.Sprintf("Error conectando la base de datos %v", err))
	}
	//	defer db.Close()

	user, err := models.ValidateUser(r)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	stmt, err := db.Prepare(`INSERT INTO users(curp, firstphone, secondphone, firstemail, secondemail, cp) VALUES ($1, $2, $3, $4, $5, $6);`)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}
	defer stmt.Close()

	// We don't need the result
	_, err = stmt.Exec(user.Curp, user.FirstPhone, user.SecondPhone, user.FirstEmail, user.SecondEmail, user.CP)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	utils.DisplayMessage(w, models.Message{
		Message: "Added! 🎉",
		Code:    http.StatusCreated,
	})

	return nil
}
