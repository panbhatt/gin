package controller

import "github.com/panbhatt/ginbooklearn/src/services"
import "net/http"

import "github.com/gin-gonic/gin"

type VideoController struct {
	service *services.VideoService
}

func GetVideoController(service *services.VideoService) *VideoController {

	return &VideoController{
		service: service,
	}

}

func (vc *VideoController) AddVideo(c *gin.Context) {
	err := vc.service.AddVideo(10)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "DONE"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DONE"})
	return
}
