package repositories

import (
	"database/sql"
	"fmt"

	"github.com/enokawa/sandbox/go/myapi/models"
)

const (
	articleNumPerPage = 5
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

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, id int) (models.Article, error) {
	const articleSql = `
		select *
		from articles
		where article_id = ?
	`

	row := db.QueryRow(articleSql, id)
	if err := row.Err(); err != nil {
		return models.Article{}, row.Err()
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, nil
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `select nice from articles where article_id = ?`
	row := tx.QueryRow(sqlGetNice, id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	var nicenum int
	if err := row.Scan(&nicenum); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	if _, err := tx.Exec(sqlUpdateNice, nicenum+1, id); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
