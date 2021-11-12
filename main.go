package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album data structure
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// some sample albums
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// responds with a list
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // alternatively, use Context.JSON
}

// adds an album from JSON received in request
func postAlbums(c *gin.Context) {
	var newAlbum album

	// use Context.BindJSON to bind JSON to album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add album to list
	albums = append(albums, newAlbum)

	// respond with new album
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// responds with an album provided an id from a parameter
func getAlbumById(c *gin.Context) {
	// retrieve the parameter id from the context
	id := c.Param("id")

	// find and return the right album
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// respond with not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// removes an album
func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	// find, delete and respond with the associated album
	for _, a := range albums {
		if a.ID == id {
			albums, a = remove(albums, a)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// respond with not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// removes an item from a slice
func remove(slice []album, a album) ([]album, album) {
	for i, s := range slice {
		if s.ID == a.ID {
			return append(slice[:i], slice[i+1:]...), a
		}
	}
	return slice, a
}

// main
func main() {
	// get the http router from Gin
	router := gin.Default()

	// assign getAlbums() as the handler for GET /albums
	router.GET("/albums", getAlbums)

	// assign postAlbums() as the handler for POST /albums
	router.POST("/albums", postAlbums)

	// assign getAlbumById() as the handler for GET /albums/:id
	router.GET("/albums/:id", getAlbumById)

	// assign deleteAlbumById() as the handler for DELETE /albums/:id
	router.DELETE("/albums/:id", deleteAlbumById)

	// start hosting
	router.Run("localhost:8080")
}
