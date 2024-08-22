package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Json type data in Gin

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year int `json:"year"`
}

var albums = []Album {
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adelle", Year: 2022},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
}

// handler (function that takes the request as a parameter)
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)

}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}