package controllers

/* Copyright (C) 2022 Rawley Fowler
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
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/repos"
)

type CommentDto struct {
	Url string
}

type BlogController struct {
	Repository *repos.BlogRepo
}

func NewBlogController(r *repos.BlogRepo) *BlogController {
	c := new(BlogController)
	c.Repository = r
	return c
}

func RegisterBlogGroup(r *gin.RouterGroup) {
	c := NewBlogController(repos.NewBlogRepo("dsn"))
	r.GET("/", c.IndexBlogPage)
	r.GET("/post/:url", c.IndividualBlogPage)
}

func (bc *BlogController) IndexBlogPage(c *gin.Context) {
	posts, err := bc.Repository.GetAllBlogPosts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error.tmpl", &gin.H{})
		return
	}
	c.HTML(http.StatusOK, "blog.tmpl", *posts)
}

func (bc *BlogController) IndividualBlogPage(c *gin.Context) {
	post, err := bc.Repository.GetBlogByUrl(c.Param("url"))
	if err != nil {
		c.HTML(http.StatusNotFound, "not_found.tmpl", &gin.H{})
		return
	}
	// TODO: Re implement captcha here. It is already included on the blog post model, though not in the database.
	// The idea is to generate a captcha for each post, and index them by title. The async service should then update each captcha every hour.
	c.HTML(http.StatusOK, "blog_post.tmpl", post)
}
