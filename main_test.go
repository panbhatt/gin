// go test -v
// go test -v -coverprofile=coverage.out
// go tool cover -html=coverage.out

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/panbhatt/ginbooklearn/src/controller"
	"github.com/panbhatt/ginbooklearn/src/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHandler(t *testing.T) {

	mockResponse := `{"message":"hello"}`

	ts := httptest.NewServer(setUpServer())
	defer ts.Close()

	fmt.Println(ts.URL)
	resp, err := http.Get(fmt.Sprintf("%s/json", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, Got %v ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Excepted status code 200, got %v ", resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	if string(responseData) != mockResponse {
		t.Fatalf("Excepted hello world message, got %v", responseData)
	}

}

func TestJsonHandler1(t *testing.T) {

	mockResponse := `{"message":"hello"}`

	ts := httptest.NewServer(setUpServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/json", ts.URL))
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	responseData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockResponse, string(responseData))

}

func SetUpLocalServer() *gin.Engine {
	return gin.Default()
}

func TestListRecipesController(t *testing.T) {
	server := SetUpLocalServer()

	server.GET("/recipes", controller.ListRecipes)
	req, _ := http.NewRequest("GET", "/recipes", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	//var recipes []models.Receipe
	var resp struct {
		Data []models.Receipe `json:"data"`
	}
	//fmt.Printf(w.Body.String())
	respJson := w.Body.String()
	fmt.Println(respJson)
	json.Unmarshal([]byte(respJson), &resp)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(resp.Data))
}
