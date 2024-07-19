package management

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 単語を追加する
func AddWord() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストボディを取得
		v := new([]model.Word)

		// リクエストボディが不正な場合はエラー
		if err := c.BindJSON(v); err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}

		// DBに単語を追加
		for _, word := range *v {
			if err := create.InsertWord(word); err != nil {
				c.JSON(500, gin.H{"error": "failed to insert word"})
				return
			}
		}

		// レスポンス
		c.JSON(200, gin.H{
			"result":  "success",
			"message": fmt.Sprintf("added %d words", len(*v)),
		})
	}
}
