package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/balajisa09/web-service-gin/pkg/models"
	"fmt"
)

func (h handler) GetAlbums(c *gin.Context){
	var albums []models.Album
	if result := h.DB.Find(&albums) ; result.Error != nil{
		fmt.Println(result.Error)
	}
	c.IndentedJSON(http.StatusOK,albums)
}
