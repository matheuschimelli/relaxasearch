package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matheuschimelli/relaxasearch/core"
)

func GETDocCount(c *gin.Context) {
	indexName := c.Param("indexName")

	index, err := core.DocCount(indexName)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}

func POSTUpserDoc(c *gin.Context) {
	indexName := c.Param("indexName")
	docId := c.Param("docId")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	index, err := core.UpsertDoc(indexName, docId, jsonData)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})

}

func DELETEDoc(c *gin.Context) {
	indexName := c.Param("indexName")
	dir, _ := os.Getwd()
	index, err := coreService.DeleteIndex(dir, indexName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}
