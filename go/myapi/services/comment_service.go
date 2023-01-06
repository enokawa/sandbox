package services

import (
	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, nil
	}

	return newComment, nil
}
