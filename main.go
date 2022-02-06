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
	"github.com/gin-gonic/gin"
	"gitlab.com/rawleyifowler/site-rework/utils"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.tmpl")
	utils.InitializeRoutes(router)
	router.Run()
}
