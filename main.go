package main

import (
	"ignite-admin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r)
}