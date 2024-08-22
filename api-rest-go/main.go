package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Json type data in Gin

type Album struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year uint `json:"year"`
}

var albums = []Album {
	{ID: 1, Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: 2, Title: "21", Artist: "Adelle", Year: 2022},
	{ID: 3, Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
}

// handler (function that takes the request as a parameter)
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)

}

func postAlbum(c *gin.Context) {
	// que ninguno de los valores sin contar ID no esten null
	
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	if newAlbum.Artist == "" || newAlbum.Title == "" || newAlbum.Year == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Ingrese todos los campos")
		return
	}

	fmt.Println(newAlbum)

	newAlbum.ID = albums[len(albums)-1].ID+1

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/postalbums", postAlbum)

	router.Run("localhost:8080")
}