package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"Director"`
	Price    string `json:"price"`
}

var movies = []movie{
	{ID: 1, Title: "Hello World1", Director: "Hello World1", Price: "1.1"},
	{ID: 2, Title: "Hello World2", Director: "Hello World2", Price: "2.2"},
	{ID: 3, Title: "Hello World3", Director: "Hello World3", Price: "3.3"},
}

func main() {
	router := gin.New()

	router.LoadHTMLGlob("/home/student/Golang-Gin-Web-Framework/Middleware/templates/*html")

	router.Use(gin.Logger(), gin.Recovery())

	router.Use(middleware1, middleware2, middleware3())

	router.GET("/movie", getAllMovies)

	auth := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"Joe":   "baseball",
		"Kelly": "1234",
	}))

	auth.GET("/movie", createMovieForm)
	auth.POST("/movie", createMovie)
	router.Run()
}

func getAllMovies(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "allmovies.html", movies)
}

func createMovieForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "createmovieform.html", nil)
}

func createMovie(ctx *gin.Context) {
	var newMovie movie
	newMovie.ID, _ = strconv.Atoi(ctx.PostForm("id"))
	newMovie.Title = ctx.PostForm("title")
	newMovie.Director = ctx.PostForm("director")
	newMovie.Price = ctx.PostForm("price")
	movies = append(movies, newMovie)
	ctx.HTML(http.StatusOK, "allmovies.html", movies)
}

func middleware1(ctx *gin.Context) {
	fmt.Println("middleware 1")
	ctx.Next()
}

func middleware2(ctx *gin.Context) {
	fmt.Println("middleware 2")
	ctx.Next()
}

func middleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware 3")
		c.Next()
	}
}
