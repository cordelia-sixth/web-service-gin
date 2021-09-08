package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// routingを扱えるようにする
	router := gin.Default()

	// GET /albums
	// getAlbums()を呼び出す
	router.GET("/albums", getAlbums)

	// GET /albums/:id
	// Ginでは、値の前にコロンを書くとpath parameterとして
	// 認識されるので Context.param("id") で取得できる
	router.GET("/albums/:id", getAlbumByID)

	// POST /albums
	// postAlbums()を呼び出す
	router.POST("/albums", postAlbums)

	// http.server 起動
	// URL localhost:8080
	router.Run("localhost:8080")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// responsするデータ
var albums = []album{
	{ID: "1", Title: "Blue Trail", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

	// 第一引数はクライアントに返すHTTP status code
	// この場合は200 ok

	// albumsをJSONに変換して返す
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
// 引数にrequestを受け取る
func postAlbums(c *gin.Context) {
	// requestされたアルバム情報を格納
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	// 受け取ったJSONをnewAlbumにbindする
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// request pathに含まれるidを抜き出してそれにマッチするアルバムを取得して
// レスポンスする
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// request pathのidを抜き出す
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			// 200 status codeとidとマッチしたアルバム情報のJSONを返す
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// 404 status codeとメッセージを返す
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
