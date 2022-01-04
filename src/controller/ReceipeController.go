package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/panbhatt/ginbooklearn/src/models"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var receipes []models.Receipe

func Init() {
	receipes = make([]models.Receipe, 0)
	file, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		fmt.Println("AN Error occured, while reading the file ", err)
	}
	_ = json.Unmarshal([]byte(file), &receipes)

}

func ListRecipes(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": receipes,
	})
}

func CreateRecipe(c *gin.Context) {
	var recipe models.Receipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()
	receipes = append(receipes, recipe)
	c.JSON(http.StatusOK, gin.H{
		"data": receipes,
	})

}

func DeleteReceipe(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i, v := range receipes {
		if v.Id == id {
			index = i
		}
	}
	fmt.Println(id, index)

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found with id " + id,
		})
		return
	}

	receipes = append(receipes[:index], receipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted. ",
	})

}

func SearchReceipe(c *gin.Context) {
	term := c.Query("tag")
	searchResult := make([]models.Receipe, 0)

	for _, v := range receipes {
		for _, tg := range v.Tags {
			if strings.EqualFold(tg, term) {
				searchResult = append(searchResult, v)
			}
		}
	}

	c.JSON(http.StatusOK, searchResult)
}

func UpdateRecipe(c *gin.Context) {
	var recipe models.Receipe
	id := c.Param("id")
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var index int = -1
	for i, v := range receipes {
		if v.Id == id {
			index = i
		}
	}

	fmt.Println("ID = ", id, " INDEX = ", index)

	if index != -1 {
		receipes[index] = recipe
		c.JSON(http.StatusOK, gin.H{
			"data": receipes,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cant foudn the id",
		})
		return
	}

}
