package repos

import (
	"testing"

	"gitlab.com/rawleyifowler/site-rework/models"
)

var cr *CommentRepo

func TestInitialize(t *testing.T) {
	var err error
	cr, err = NewCommentRepo("dsn")
	if err != nil {

	}
}

func TestGetCommentById(t *testing.T) {
	const ID uint = 1
	const AUTHOR string = "Whoops"
	const CONTENT string = "Uhhhhhhhh that wasn't supposed to happen..."
	// This comment will exist forever on the database, it's part of the test blog post.
	x := models.Comment{Id: ID, Author: AUTHOR, Content: CONTENT}
	y, err := cr.GetCommentById(ID)
	if err != nil {
		t.Fatalf("Expected: Comment with id: %d should be found without errors.", ID)
	}
	if x.Id != y.Id ||
		x.Author != y.Author ||
		x.Content != y.Content {
		t.Fatalf("Expected(left expected, right result):\nid -> %d == %d\nauthor -> %s == %s\ncontent -> %s == %s",
			ID, y.Id, AUTHOR, y.Author, CONTENT, y.Content)
	}
	_, err := cr.GetCommentById(0)
	if err == nil {
		t.Fatalf("Expected: Values of less than or equal to 0 are considered illegal values for comment id")
	}
}

func TestGetCommentsByAssociatedPost(t *testing.T) {
	_, err := cr.GetCommentsByAssociatedPost("")
	if err == nil {
		t.Fatalf("Expected: Empty values are considered illegal for associated post")
	}
	// Wild associated post should return empty array
	post, err := cr.GetCommentsByAssociatedPost("kjdoJDOWIJWKIJALSKAJSAKIJDOIAj")
	if len(*post) > 0 {
		t.Fatalf("Expected: Should return empty array")
	}
	if err != nil {
		t.Fatalf("Expected: Non existant post should not throw error, just empty array")
	}
}

func TestGetCommentsByAuthor(t *testing.T) {
	_, err := cr.GetCommentsByAuthor("")
	if err == nil {
		t.Fatalf("Expected: An empty author value should throw an error")
	}
}

func TestCreateComment(t *testing.T) {
	err := cr.CreateComment(nil)
	if err == nil {
		t.Fatalf("Expected: Nil values should result in errors when trying to create comments")
	}
	err = cr.CreateComment(&models.Comment{})
	if err == nil {
		t.Fatalf("Expected: Empty or incomplete models should result in errors.")
	}
	err = cr.CreateComment(&models.Comment{Author: "Chuck"})
	if err == nil {
		t.Fatalf("Expected: Empty or incomplete models should result in errors.")
	}
}
