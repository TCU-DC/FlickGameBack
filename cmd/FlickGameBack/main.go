package main

import (
	"FlickGameBack/pkg/db/create"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのテーブルを作成
	create.CreateDefaultTable()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "FlickGameBack",
		})
	})
	r.Run(":8080")
}
