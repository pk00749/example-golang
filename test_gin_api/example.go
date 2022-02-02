package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John", Price: 56.99},
	{ID: "2", Title: "Yellow", Artist: "Jack", Price: 55.99},
	{ID: "3", Title: "Red", Artist: "Jay", Price: 54.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusFound,
		gin.H{"message": "album not found"},
	)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/postAlbums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
