package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/enokawa/sandbox/go/myapi/controllers"
	"github.com/enokawa/sandbox/go/myapi/services"
	"github.com/gorilla/mux"
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
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := mux.NewRouter()

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
