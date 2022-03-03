package main

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
	"html/template"

	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/bootstrap"
)

var router *gin.Engine
var dsn string

func main() {
	router = gin.Default()
	router.Use(gin.Recovery())
	// Create function map for templates that will unescape HTML from the database
	// Credits to: https://h1z3y3.me/posts/go-html-template-script-unescape
	funcMap := make(template.FuncMap)
	funcMap["UnescapeHTML"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	router.SetFuncMap(funcMap)
	// This needs to change because of RCCTL in OpenBSD, not sure if you can use a ksh variable as path in Gin but i'll test
	router.LoadHTMLGlob("templates/*.tmpl")
	router.TrustedPlatform = "X-Real-Ip"
	bootstrap.InitializeRoutes(router)
	router.Run(":8080")
}
