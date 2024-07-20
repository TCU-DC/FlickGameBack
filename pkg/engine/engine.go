package engine

import (
	"FlickGameBack/pkg/engine/core"
	"FlickGameBack/pkg/engine/management"
	"FlickGameBack/pkg/engine/ranking"
	"FlickGameBack/pkg/engine/score"
	"FlickGameBack/pkg/engine/socket"

	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {

	// ルーティング
	r.GET("/word-get", core.WordGet())
	r.POST("/add-words", management.AddWord())
	r.POST("/register-score", score.RegisterScore())
	r.GET("/get-ranking", ranking.GetRanking())

	r.GET("/room-add", socket.AddRoomSession)
	r.GET("/room-join", socket.JoinRoomSession)
	// ソケット通信エンドポイント
	r.GET("/room/:id", socket.HandleWebSocket)
	return r
}
