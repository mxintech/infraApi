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
	db := db.GetDatabase(w)

	user, err := models.ValidateUser(r)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusBadRequest,
		})
		return
	}

	stmt, err := db.Prepare(`INSERT INTO users 
							VALUES (?, ?, ?, ?, ?, ?);`)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	defer stmt.Close()

	results, err := stmt.Exec(user.Curp, user.FirstPhone, user.SecondPhone, user.FirstEmail, user.SecondEmail, user.CP)
	if err != nil {
		utils.DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
	}

	utils.DisplayMessage(w, models.Message{
		Message: fmt.Sprintf("%v", results),
		Code:    http.StatusCreated,
	})
}
