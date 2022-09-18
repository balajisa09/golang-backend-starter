package handlers 

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/balajisa09/web-service-gin/pkg/models"
)

func (h handler) PostAlbums(c *gin.Context){
	var newalbum models.Album
	if err:= c.BindJSON(&newalbum); err !=nil{
		return
	}
	
	h.DB.Create(newalbum)
	c.IndentedJSON(http.StatusCreated,newalbum)
}
