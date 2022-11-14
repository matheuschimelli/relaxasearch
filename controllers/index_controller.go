package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matheuschimelli/relaxasearch/core"
)

type Response struct {
	Success bool
	Message string
}

var coreService = core.NewIndexService()

func GETAllIndexes(c *gin.Context) {

	data := coreService.ListIndexes()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func GETShowIndex(c *gin.Context) {
	indexName := c.Param("indexName")

	index, err := coreService.ShowIndex(indexName)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}

func POSTCreateIndex(c *gin.Context) {
	indexName := c.Param("indexName")
	dir, _ := os.Getwd()
	index, err := coreService.CreateIndex(dir, indexName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}

func DELETEIndex(c *gin.Context) {
	indexName := c.Param("indexName")
	dir, _ := os.Getwd()
	index, err := coreService.DeleteIndex(dir, indexName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}
