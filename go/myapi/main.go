package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/enokawa/sandbox/go/myapi/handlers"
	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/gorilla/mux"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	article_id := 1
	const sqlGetNice = `select nice from articles where article_id = ?;`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var article models.Article
	err = row.Scan(&article.NiceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	article.NiceNum++

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?;`
	if _, err := tx.Exec(sqlUpdateNice, article.NiceNum, article_id); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
