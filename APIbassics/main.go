package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Price    string `json:"price"`
}

var movies = []movie{
	{ID: 1, Title: "Hello World1", Director: "Hello World1", Price: "1.1"},
	{ID: 2, Title: "Hello World2", Director: "Hello World2", Price: "2.2"},
	{ID: 3, Title: "Hello World3", Director: "Hello World3", Price: "3.3"},
}

func main() {
	router := gin.Default()

	router.GET("/movie", getMovie)
	router.GET("/movie/:id", getMovieByID)
	router.POST("/movie", createPost)
	router.PATCH("/movie/:id", updateMoviePrice)
	router.DELETE("/movie/:id", deleateMovie)

	router.Run()
}

func getMovie(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, movies)
}

func getMovieByID(ctx *gin.Context) {
	var index int
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	for i, v := range movies {
		if id == v.ID {
			index = i
		}
	}
	ctx.JSON(http.StatusOK, movies[index])
}

// createPost

func createPost(ctx *gin.Context) {
	var newMovie movie
	err := ctx.BindJSON(&newMovie)
	if err != nil {
		log.Println("err", err)
		return
	}
	movies = append(movies, newMovie)
	ctx.JSON(http.StatusOK, movies)
}

func updateMoviePrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalln(err, "update movie price")
		return
	}
	var index int

	for i, v := range movies {
		if id == v.ID {
			index = i
		}
	}
	movies[index].Price = "10.10"
	ctx.JSON(http.StatusOK, movies)
}

func deleateMovie(ctx *gin.Context) {
	var index int
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("error delete movie", err)
		return
	}
	for i, v := range movies {
		if id == v.ID {
			index = i
		}
	}
	movies = append(movies[:index], movies[index+1:]...)
	ctx.JSON(http.StatusOK, movies)
}
