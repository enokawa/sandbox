package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/enokawa/sandbox/go/myapi/api"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("USERNAME")
	dbPassword = os.Getenv("USERPASS")
	dbDatabase = os.Getenv("DATABASE")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
