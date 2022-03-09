package repos

/*
Copyright (C) 2022 Rawley Fowler

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.Rawley Fowler, 2022
*/

import (
	"testing"

	"gitlab.com/rawleyifowler/site-rework/models"
)

var test_blog_repo = NewBlogRepo("dsn")

func TestGetBlogByUrl(t *testing.T) {
	_, err := test_blog_repo.GetBlogByUrl("")
	if err == nil {
		t.Fatalf("Expected: Empty values of url for get blog by url should return an error")
	}
}

func TestCreateBlogPost(t *testing.T) {
	err := test_blog_repo.CreateBlogPost(&models.BlogPost{})
	if err == nil {
		t.Fatalf("Expected: Blog posts must be fully formed")
	}

	err = test_blog_repo.CreateBlogPost(&models.BlogPost{Title: "abc"})
	if err == nil {
		t.Fatalf("Expected: Blog posts must be fully formed, title, content, url")
	}

	err = test_blog_repo.CreateBlogPost(nil)
	if err == nil {
		t.Fatalf("Expected: Cannot create blog posts from nil values")
	}
}
