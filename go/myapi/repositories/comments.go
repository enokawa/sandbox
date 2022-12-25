package repositories

import (
	"database/sql"
	"fmt"

	"github.com/enokawa/sandbox/go/myapi/models"
)

func GetComment(db *sql.DB, id int) ([]models.Comment, error) {
	const sqlStr = `
		select commend_id, article_id, message, created_at
		from comments
		where article_id = ?
	`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Println(err)
		return []models.Comment{}, nil
	}
	defer db.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
