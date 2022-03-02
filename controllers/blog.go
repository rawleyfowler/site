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
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CommentDto struct {
	Url string
}

var (
	db            *gorm.DB
	recentPosters map[string]uint
	captchaVals   [2]int
)

func RegisterBlogGroup(r *gin.RouterGroup) {
	// Initialize recent posters cache
	recentPosters = make(map[string]uint)
	// Load dsn and initialize database
	dsn := utils.LoadDSN("dsn")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed...")
	}
	utils.PerformMigrations(db)
	r.GET("/", RenderBlogPage)
	r.GET("/post/:url", RenderIndividualBlogPost)
	r.POST("/post")
	r.POST("/post/comment", CreateComment)
	// Clear the commenters cache every 15 minutes
	go utils.TimedClearMap(&recentPosters, 180000*5)
	go utils.UpdateCaptcha(&captchaVals)
}

func RenderBlogPage(c *gin.Context) {
	posts := GetAllBlogPosts()
	if posts == nil {
		c.HTML(http.StatusInternalServerError, "internal_server_error.tmpl", models.Page{Title: " | 500"})
	} else {
		c.HTML(http.StatusOK, "blog.tmpl", struct {
			Posts []models.BlogPost
			Title string
		}{Posts: *posts, Title: " | blog"})
	}
}

func RenderIndividualBlogPost(c *gin.Context) {
	bp := GetBlogPostById(c.Param("url"))
	if bp == nil {
		c.HTML(http.StatusNotFound, "not_found.tmpl", models.Page{Title: "404"})
	} else {
		// Blog posts handle the title themselves
		c.HTML(http.StatusOK, "blog_post.tmpl", struct {
			Post    *models.BlogPost
			Captcha [2]int
		}{Post: bp, Captcha: captchaVals})
	}
}

func CreateComment(c *gin.Context) {
	if c.Request.ParseForm() != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	a := []string{c.Request.Form.Get("author"),
		c.Request.Form.Get("content"),
		c.Request.Form.Get("url"),
		c.Request.Form.Get("captcha")}
	// If the length of the comment is great enough, and the comment already exists we can safely assume it is spam.
	if len(a[1]) > 20 && len(*GetCommentsByContent(a[1], a[2])) > 0 {
		c.HTML(http.StatusNotAcceptable, "comment_post_failed.tmpl", CommentDto{Url: a[2]})
		c.Abort()
	}
	i, err := strconv.ParseInt(a[3], 10, 32)
	if err != nil || int(i) != (captchaVals[0]+captchaVals[1]) {
		c.HTML(http.StatusNotAcceptable, "comment_post_failed.tmpl", CommentDto{Url: a[2]})
		return
	}
	if GetNumberOfRecentPosts(c) > 3 {
		c.HTML(http.StatusNotAcceptable, "comment_post_spam.tmpl", CommentDto{Url: a[2]})
		return
	}
	for _, v := range a {
		if len(v) == 0 {
			c.HTML(http.StatusNotAcceptable, "comment_post_failed.tmpl", CommentDto{Url: a[2]})
			return
		}
	}
	comment := models.Comment{
		Author:         a[0],
		Content:        a[1],
		AssociatedPost: a[2],
	}
	db.Create(&comment)
	recentPosters[c.ClientIP()]++
	// Pass the associated post to the template to add to the href
	c.HTML(http.StatusOK, "comment_post.tmpl", CommentDto{Url: comment.AssociatedPost})
}

func GetNumberOfRecentPosts(c *gin.Context) uint {
	return recentPosters[c.ClientIP()]
}

func AddBlogPost(bp *models.BlogPost) bool {
	err := db.Create(bp).Error
	return err == nil
}

func DeleteBlogPostById(id string) bool {
	err := db.Model(&models.BlogPost{}).Delete(&models.BlogPost{Url: id}).Error
	return err == nil
}

func GetAllBlogPosts() *[]models.BlogPost {
	var posts []models.BlogPost
	// Select title, date, and url fields from the blog post records an store them in posts.
	// This is so we don't grab the entire blog post when we render them on the overview page. Saves a couple ms.
	err := db.Model(&models.BlogPost{}).Select("title, date, url").Scan(&posts).Error
	if err != nil {
		return nil
	}
	return &posts
}

func GetBlogPostById(id string) *models.BlogPost {
	var post models.BlogPost
	err := db.Model(&post).Where(&models.BlogPost{Url: id}).Find(&post).Error
	if err != nil {
		return nil
	}
	// TODO: Figure out gorm joins!
	// gorm joins are not working at all, just going to do another query to make it work for now.
	err = db.Model(&models.Comment{}).Where(&models.Comment{AssociatedPost: id}).Find(&post.Comments).Error
	if err != nil {
		return nil
	}
	return &post
}

func GetCommentsByContent(content string, url string) *[]models.Comment {
	// This is used to make sure people don't paste the same thing over and over.
	var comments []models.Comment
	err := db.Model(&comments).Where("content like ? and associated_post like ?", "%"+content[1:]+"%", url).Scan(&comments).Error
	if err != nil {
		return &[]models.Comment{}
	}
	return &comments
}
