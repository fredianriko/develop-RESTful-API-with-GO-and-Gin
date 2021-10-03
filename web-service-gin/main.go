package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func postAlbums(c *gin.Context){
// 	var newAlbum album
// 	//Call BindJSON to bind the received JSON to newAlbum

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	//add the new album to the slice
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

//make json structure and it's data type
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
// initialize the album slice that going to be turn into json
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Colrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan & Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
// the main function that route getAlbums function, and determine the endpoint api to get json
func main() {
	router := gin.Default()
	router.GET("/albums	", getAlbums) //get the albums slice, and passed it to albums, the endpoint has to be the slice name
	// router.POST("/albums ", postAlbums)
	router.Run("localhost:8080")// root api at localhost:8080 , then the api become localhost:8080/albums to get the json of albums
}

//function to turn albums slice into json format
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


// run these command after type all these code
// go get .
// go run .
// open new terminal and type "localhost:8080/albums"
// or open browser and type "localhost:8080/albums" to see the full json