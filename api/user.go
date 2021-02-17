package api

import (
	"fmt"
	"net/http"

	"github.com/TheGolurk/infraApi/db"
	"github.com/TheGolurk/infraApi/models"
	"github.com/TheGolurk/infraApi/utils"

	_ "github.com/lib/pq" // Postgres Driver
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDatabase()
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%s %v", "Error conectando la base de datos", err),
			Code:    http.StatusBadRequest,
		})
		return
	}
	defer db.Close()

	user, err := models.ValidateUser(r)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusBadRequest,
		})
		return
	}

	stmt, err := db.Prepare(`INSERT INTO users(curp, firstphone, secondphone, firstemail, secondemail, cp) VALUES ($1, $2, $3, $4, $5, $6);`)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	defer stmt.Close()

	// We don't need the result
	_, err = stmt.Exec(user.Curp, user.FirstPhone, user.SecondPhone, user.FirstEmail, user.SecondEmail, user.CP)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
	}

	utils.DisplayMessage(w, models.Message{
		Message: "Added! 🎉",
		Code:    http.StatusCreated,
	})
}
