package repositories

import (
	"database/sql"
	"fmt"

	"github.com/enokawa/sandbox/go/myapi/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, ?, now());
	`

	_, err := db.Exec(sqlStr, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
	if err != nil {
		fmt.Println(err)
		return article, nil
	}

	return article, nil
}
