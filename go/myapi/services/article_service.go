package services

import (
	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// TODO: sql.DB 型を手に入れて、変数 db に代入する
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}
