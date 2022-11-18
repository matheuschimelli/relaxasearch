package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/matheuschimelli/relaxasearch/controllers"
	core "github.com/matheuschimelli/relaxasearch/core"
)

//var dataDir = flag.String("dataDir", "data", "data directory")

func main() {
	router := gin.Default()

	coreService := core.NewIndexService()
	coreService.InitIndex("relaxasearchData")

	// Handles index
	router.GET("/relaxasearch/v1", controller.GETAllIndexes)
	router.GET("/relaxasearch/v1/:indexName", controller.GETShowIndex)
	router.POST("/relaxasearch/v1/:indexName", controller.POSTCreateIndex)
	router.DELETE("/relaxasearch/v1/:indexName", controller.DELETEIndex)

	// Handles Index data
	router.GET("/relaxasearch/v1/:indexName/:docId", controller.GETShowDoc)
	router.POST("/relaxasearch/v1/:indexName/:docId", controller.POSTUpserDoc)

	router.POST("/relaxasearch/v1/:indexName/search", controller.POSTSearch)
	router.GET("/relaxasearch/v1/:indexName/count", controller.GETDocCount)
	router.DELETE("/relaxasearch/v1/:indexName/:docId", controller.DELETEDoc)

	router.Run(":3000")
}
