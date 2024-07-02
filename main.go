package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `form:"id"`
	Title  string  `form:"title"`
	Artist string  `form:"artist"`
	Price  float64 `form:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.Use(Logger())

	albums := router.Group("/albums")
	{
		albums.GET("", getAlbums)
		albums.GET(":id", getAlbumByID)
		albums.POST("", postAlbums)
	}

	router.Run("localhost:8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("latencia:", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBind(&newAlbum); err != nil {
		// show error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
