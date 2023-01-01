package repositories_test

import (
	"testing"

	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"
	"github.com/enokawa/sandbox/go/myapi/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "enokawa",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	article := models.Article{
		Title:    "UpdateTest",
		Contents: "testtest",
		UserName: "enokawa",
	}
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	if err := repositories.UpdateNiceNum(testDB, newArticle.ID); err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, newArticle.ID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum != 1 {
		t.Errorf("fail to update nice num")
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, newArticle.Title, newArticle.Contents, newArticle.UserName)
	})
}
