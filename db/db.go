package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Postgres Drive
)

var (
	dbname = os.Getenv("db")
	pwd    = os.Getenv("pwd")
	host   = os.Getenv("host")
	user   = os.Getenv("user")
	port   = os.Getenv("port")
	url    = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=require", host, port, dbname, user, pwd)
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
