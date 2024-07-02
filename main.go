package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leidison/golang-simple-rest/middlewares"
)

type album struct {
	ID     string  `form:"id" binding:"required"`
	Title  string  `form:"title" binding:"required"`
	Artist string  `form:"artist"`
	Price  float64 `form:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	godotenv.Load()

	router := gin.Default()

	// middlewares
	jwtMiddleware := middlewares.Jwt()

	// global middlewares
	router.Use(middlewares.Logger())

	// routes

	albums := router.Group("/albums")
	{
		albums.GET("", getAlbums)
		albums.GET(":id", getAlbumByID)

		private := albums.Group("/").Use(jwtMiddleware)
		{
			private.POST("", postAlbums)
		}
	}

	router.Run("localhost:3000")
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var input album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// Add the new album to the slice.
	albums = append(albums, input)

	c.JSON(http.StatusCreated, input)
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
