package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/enokawa/sandbox/go/myapi/models"
	"github.com/enokawa/sandbox/go/myapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "enokawa",
		NiceNum:  4,
	}

	got, err := repositories.SelectArticleDetail(db, 1)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Contents: get %s but want %s", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d", got.NiceNum, expected.NiceNum)
	}
}
