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
	"bytes"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/rawleyfowler/rawleydotxyz/repos"
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
	// RSS Feed
	r.GET("/feed", c.RSSFeed)
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
	// Using the repository, get the plug that correlates to the url
	post, err := bc.Repository.GetBlogByUrl(c.Param("url"))
	if err != nil {
		c.HTML(http.StatusNotFound, "not_found.tmpl", &gin.H{})
		return
	}
	c.HTML(http.StatusOK, "blog_post.tmpl", post)
}

func (bc *BlogController) RSSFeed(c *gin.Context) {
	posts, err := bc.Repository.GetAllBlogPosts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error.tmpl", &gin.H{})
		return
	}
	t, err := template.ParseFiles("templates/rss.tmpl")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error", &gin.H{})
		return
	}
	c.Header("Content-Type", "text/xml")
	var rawXml bytes.Buffer
	err = t.Execute(&rawXml, posts)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error", &gin.H{})
		return
	}
	c.String(http.StatusOK, strings.Replace(rawXml.String(), "&", "&amp;", -1), &gin.H{})
}
