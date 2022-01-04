package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/panbhatt/ginbooklearn/src/controller/"
	"github.com/panbhatt/ginbooklearn/src/controller"
	"github.com/panbhatt/ginbooklearn/src/middlewares"
	"github.com/panbhatt/ginbooklearn/src/services"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	server := setUpServer()
	server.Run(":5050")

}

func setUpServer() *gin.Engine {
	controller.Init()

	var videoService *services.VideoService = services.GetVideoService()
	var videoController *controller.VideoController = controller.GetVideoController(videoService)
	var authController *controller.AuthController = &controller.AuthController{}

	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/json", JsonHandler)
	router.GET("/prometheus", gin.WrapH(promhttp.Handler()))
	router.Static("/assets", "./assets")
	router.GET("/recipes", controller.ListRecipes)
	router.POST("/recipes", controller.CreateRecipe)
	router.PUT("/recipes/:id", controller.UpdateRecipe)
	router.DELETE("/recipes/:id", controller.DeleteReceipe)
	router.GET("/recipes/search", controller.SearchReceipe)

	apiRouter := router.Group("/api")
	apiRouter.Use(middlewares.AuthMiddleware())
	apiRouter.GET("/video", videoController.AddVideo)

	router.GET("/:name", PathNameHandler)

	router.POST("/jwt/token", authController.VerifyUser)

	return router
}

type Person struct {
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lasttName,attr"`
}

func IndexHandler(c *gin.Context) {
	c.File("index.html")
}

func JsonHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func PathNameHandler(c *gin.Context) {

	//name := c.Params.ByName("name")
	// c.JSON(200, gin.H{
	// 	"message": "HEllo " + name,
	// })
	c.XML(200, Person{"pankaj", "bhatt"})

}
