package repositories_test

import (
	"testing"

	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"
)

func TestSelectComment(t *testing.T) {
	article := models.Article{
		Title:    "SelectCommentTest",
		Contents: "SelectCommentTest",
		UserName: "enokawa",
	}
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	before := models.Comment{
		ArticleID: newArticle.ID,
		Message:   "SelectCommentTest",
	}
	newComment, err := repositories.InsertComment(testDB, before)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectComment(testDB, newComment.CommentID)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range after {
		if before.Message != v.Message {
			t.Errorf("new comment message is expected %s but got %s\n", before.Message, v.Message)
		}
	}

	t.Cleanup(func() {
		const commentSql = `delete from comments where comment_id = ?`
		testDB.Exec(commentSql, newComment.CommentID)

		const articleSql = `delete from articles where article_id = ?`
		testDB.Exec(articleSql, newArticle.ID)
	})
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
