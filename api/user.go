package api

import (
	"fmt"
	"net/http"

	"github.com/TheGolurk/infraApi/db"
	"github.com/TheGolurk/infraApi/models"

	_ "github.com/lib/pq" // Postgres Driver
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetDatabase()

	user, err := models.ValidateUser(r)
	if err != nil {
		DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusBadRequest,
		})
		return
	}

	query, err := db.Prepare("INSERT INTO users(curp, firstphone, secondphone, firstemail, secondemail, cp) VALUESI(?, ?, ?, ?, ?, ?)")
	if err != nil {

	}
	defer query.Close()

	results, err := query.Exec(user.CP, user.FirstPhone, user.SecondPhone, user.FirstEmail, user.SecondEmail, user.CP)
	if err != nil {
		DisplayMessage(w, models.Message{
			Message: fmt.Sprintf("%v", err),
			Code:    http.StatusInternalServerError,
		})
	}

	DisplayMessage(w, models.Message{
		Message: fmt.Sprintf("%v", results),
		Code:    http.StatusCreated,
	})
}
