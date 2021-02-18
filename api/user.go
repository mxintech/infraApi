package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/TheGolurk/infraApi/models"
	"github.com/TheGolurk/infraApi/utils"

	_ "github.com/lib/pq" // Postgres Driver
)

func CreateUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) error {
	//	defer db.Close()

	var user models.User
	err := models.ValidateUser(r, user)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = conn.ExecContext(ctx, `INSERT INTO users(curp, firstphone, secondphone, firstemail, secondemail, cp) 
	VALUES ($1, $2, $3, $4, $5, $6);`, user.Curp, user.FirstPhone, user.SecondPhone, user.FirstEmail, user.SecondEmail, user.CP)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	utils.DisplayMessage(w, models.Message{
		Message: "Added! 🎉",
		Code:    http.StatusCreated,
	})

	fmt.Println(conn.Stats().OpenConnections)
	return nil
}
