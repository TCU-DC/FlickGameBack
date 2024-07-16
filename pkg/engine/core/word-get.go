package core

import (
	"github.com/gin-gonic/gin"
)

// 単語を取得するエンドポイントハンドラ
func WordGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		word := "apple"
		c.JSON(200, gin.H{
			"word": word,
		})
	}
}
