package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello from controller!"})
}

func GetAlbums(c *gin.Context) {
	// c.JSON(http.StatusOK, albums)
	c.JSON(http.StatusOK, gin.H{"data": "hello from controller!"})
}
