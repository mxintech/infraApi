package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" // Postgres Drive
)

var (
	dbname  = os.Getenv("db")
	pwd     = os.Getenv("pswd")
	host    = os.Getenv("host")
	user    = os.Getenv("user")
	port, _ = strconv.Atoi(os.Getenv("port"))
	url     = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=disable", host, port, dbname, user, pwd)
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
