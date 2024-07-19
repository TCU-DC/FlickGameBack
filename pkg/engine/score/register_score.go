package score

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/model"
	"FlickGameBack/pkg/util"

	"github.com/gin-gonic/gin"
)

func RegisterScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var score model.Score
		// リクエストボディをパース
		c.BindJSON(&score)

		// もしユーザIDが空文字ならゲストとして登録(nullも判定可能)
		if score.UserID == "" {
			score.UserID = "guest_" + util.RandomString(10)
		}

		// 得点情報を登録
		err := create.InsertScore(score.UserID, score.Point, score.AverageSpeed, score.Level)
		if err != nil {
			c.JSON(500, gin.H{"message": "failed to insert score"})
			return
		}

		// レスポンス
		c.JSON(200, gin.H{"message": "success"})
	}
}
