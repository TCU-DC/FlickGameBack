package ranking

import (
	"FlickGameBack/pkg/db/read"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ランキングを取得するエンドポイントハンドラ
func GetRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		// クエリパラメータを取得
		level := c.Query("level")
		count_str := c.Query("high_order")

		// クエリパラメータがない場合はエラー
		if level == "" || count_str == "" {
			c.JSON(400, gin.H{"error": "level and count are required"})
			return
		}

		// high_orderパラメータをintに変換
		count, err := strconv.Atoi(count_str)
		if err != nil {
			c.JSON(400, gin.H{"error": "high_order must be an integer"})
			return
		}

		// DBからランキングを取得
		ranking, err := read.ReadRanking(level, count)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read ranking"})
			return
		}

		// レスポンス
		c.JSON(200, gin.H{
			"level":   level,
			"ranking": ranking,
		})
	}
}
