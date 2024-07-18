package core

import (
	"FlickGameBack/pkg/db/read"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 単語を取得するエンドポイントハンドラ
func WordGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		// クエリパラメータを取得
		level := c.Query("level")
		count_str := c.Query("count")

		// クエリパラメータがない場合はエラー
		if level == "" || count_str == "" {
			c.JSON(400, gin.H{"error": "level and count are required"})
			return
		}

		// countパラメータをintに変換
		count, err := strconv.Atoi(count_str)
		if err != nil {
			c.JSON(400, gin.H{"error": "count must be an integer"})
			return
		}

		// DBから単語を取得
		words, err := read.ReadWords(level, count)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read words"})
			return
		}

		// レスポンス
		c.JSON(200, gin.H{"data": words})
	}
}
