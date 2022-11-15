package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

func GETShowDoc(c *gin.Context) {
	indexName := c.Param("indexName")
	docId := c.Param("docId")

	index, err := core.ShowDoc(indexName, docId)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}

func POSTUpserDoc(c *gin.Context) {
	indexName := c.Param("indexName")

	jsonSearchQuery, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	index, err := core.Search(indexName, jsonSearchQuery)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})

}

func POSTSearch(c *gin.Context) {
	indexName := c.Param("indexName")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	index, err := core.Search(indexName, jsonData)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})

}

func DELETEDoc(c *gin.Context) {
	indexName := c.Param("indexName")
	docId := c.Param("docId")

	index, err := core.DeleteDoc(indexName, docId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": index})
}
