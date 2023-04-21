package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/hello", getHello)
	router.GET("/greet", getGreet)
	router.GET("/greet/:name", getGreetName)
	router.GET("/many", getManyData)
	router.GET("/form", getForm)
	router.POST("/form", postForm)

	router.Run()
}

func getHello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World")
}

func getGreet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "greting.html", nil)
}

func getGreetName(ctx *gin.Context) {
	name1, ok := ctx.Copy().Params.Get("name")
	if !ok {
		log.Fatal(name1, ok)
		return
	}

	name := ctx.Param("name")
	log.Println(name)
	ctx.HTML(http.StatusOK, "customGreeting.html", name)
}

func getManyData(ctx *gin.Context) {
	foods := []string{"chicken sandwich", "fries", "soda", "cookie"}

	ctx.HTML(http.StatusOK, "manyData.html", gin.H{
		"name": "Carl",
		"food": foods,
	})
}

func getForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form.html", nil)
}

func postForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	food := ctx.PostForm("food")
	ctx.HTML(http.StatusOK, "formResult.html", gin.H{
		"name": name,
		"food": food,
	})
}
