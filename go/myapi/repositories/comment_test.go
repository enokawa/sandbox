package repositories_test

import (
	"testing"

	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleId := 1
	got, err := repositories.SelectCommentList(testDB, articleId)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleId {
			t.Errorf("new comment message is expected %d but got %d\n", articleId, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "InsertCommentTest",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `delete from comments where message = ?`
		testDB.Exec(sqlStr, comment.Message)
	})
}
