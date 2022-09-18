package main

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/balajisa09/web-service-gin/pkg/handlers"
	db "github.com/balajisa09/web-service-gin/pkg/db"
)

func main(){
	DB := db.Init()
	h := handlers.New(DB)

	router := gin.Default()

	router.GET("/albums",h.GetAlbums)
	router.POST("/albums",h.PostAlbums)
	router.GET("/albums/:id", h.GetAlbumByID)
	router.Run("localhost:8080")
}
