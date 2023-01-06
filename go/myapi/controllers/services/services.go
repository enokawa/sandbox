package services

import "github.com/enokawa/sandbox/go/myapi/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(coment models.Comment) (models.Comment, error)
}
