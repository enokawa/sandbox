package testdata

import "github.com/enokawa/sandbox/go/myapi/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "enokawa",
		NiceNum:  4,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "enokawa",
		NiceNum:  4,
	},
}
