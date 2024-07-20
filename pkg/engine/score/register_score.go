package score

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/model"
	"FlickGameBack/pkg/util"

	"github.com/gin-gonic/gin"
)

func RegisterScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req_score model.RequestScore
		// リクエストボディをパース
		c.BindJSON(&req_score)

		// もしユーザIDが空文字ならゲストとして登録(nullも判定可能)
		if req_score.UserID == "" {
			req_score.UserID = "guest_" + util.RandomString(10)

			// ゲスト名が空文字ならゲスト名を設定
			if req_score.GuestName == "" {
				req_score.GuestName = "名無しさん"
			}

			// ゲストの場合はゲストユーザを登録
			err := create.InsertUser(model.User{
				UserID:        req_score.UserID,
				UserName:      req_score.UserID,
				NickName:      req_score.GuestName,
				Password_hash: "",
			})
			if err != nil {
				c.JSON(500, gin.H{"message": "failed to insert user"})
				return
			}
		}

		// 得点情報を登録
		err := create.InsertScore(req_score.UserID, req_score.Point, req_score.AverageSpeed, req_score.Level)
		if err != nil {
			c.JSON(500, gin.H{"message": "failed to insert req_score"})
			return
		}

		// レスポンス
		c.JSON(200, gin.H{"message": "success"})
	}
}
