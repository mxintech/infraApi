package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq" // Postgres Drive
)

var (
	dbname  = os.Getenv("db")
	pwd     = os.Getenv("pswd")
	host    = os.Getenv("host")
	user    = os.Getenv("user")
	port, _ = strconv.Atoi(os.Getenv("port"))
	url     = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=verify-full", host, port, dbname, user, pwd)
	conn    *sql.DB
	once    sync.Once
)

func GetDatabase() (*sql.DB, error) {
	fmt.Println(url)

	var err error
	once.Do(func() {
		conn, err = sql.Open("postgres", url)
		if err != nil {
			log.Panic(err)
		}

		if err = conn.Ping(); err != nil {
			log.Panic(err)
		}

		conn.SetMaxOpenConns(0)
		conn.SetMaxIdleConns(0)
		conn.SetConnMaxLifetime(time.Nanosecond)

	})

	return conn, err
}
