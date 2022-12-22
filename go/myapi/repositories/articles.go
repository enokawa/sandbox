package repositories

import (
	"database/sql"
	"fmt"

	"github.com/enokawa/sandbox/go/myapi/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, nil
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func ListArticle(db *sql.DB, page int) ([]models.Article, error) {
	return []models.Article{}, nil
}
