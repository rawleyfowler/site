package controllers

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
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.com/rawleyifowler/site-rework/models"
	"gitlab.com/rawleyifowler/site-rework/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterBlogGroup(r *gin.RouterGroup) {
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
}

func RenderBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.tmpl", *GetAllBlogPosts())
}

func RenderIndividualBlogPost(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_post.tmpl", GetBlogPostById(c.Param("url")))
}

func GetAllBlogPosts() *[]models.BlogPost {
	var posts []models.BlogPost
	db.Find(&posts)
	return &posts
}

func GetBlogPostById(id string) *models.BlogPost {
	var post models.BlogPost
	db.Where(&models.BlogPost{ Url: id }).First(&post)
	return &post
}
