package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/balajisa09/web-service-gin/pkg/models"
    "fmt"
)

func (h handler) GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    var albums []models.Album
	if result := h.DB.Find(&albums) ; result.Error != nil{
		fmt.Println(result.Error)
	}

    for _, a := range albums{
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}