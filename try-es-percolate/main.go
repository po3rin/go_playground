package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func main() {
	r := gin.Default()
	r.GET("/:id", func(c *gin.Context) {
		url := "http://elasticsearch:9200"
		client, err := elastic.NewClient(
			elastic.SetURL(url),
			elastic.SetSniff(false),
			elastic.SetBasicAuth("elastic", "changeme"),
		)
		if err != nil {
			c.JSON(500, gin.H{
				"fase":    "init client",
				"message": err.Error(),
			})
			return
		}
		defer client.Stop()

		id := c.Param("id")
		percolateClient := elastic.NewPercolatorQuery()
		query := percolateClient.Field("query").
			DocumentType("articles").
			IndexedDocumentType("title").
			IndexedDocumentIndex("testindex").
			IndexedDocumentId(id)

		result, err := client.Search().
			Index("testindex").
			Query(query).
			Do(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{
				"fase":    "search to elasticsearch",
				"message": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"message": result,
		})
	})
	r.Run()
}
