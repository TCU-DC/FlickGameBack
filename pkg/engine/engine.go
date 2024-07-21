package engine

import (
	"FlickGameBack/pkg/engine/core"
	"FlickGameBack/pkg/engine/management"
	"FlickGameBack/pkg/engine/ranking"
	"FlickGameBack/pkg/engine/score"
	"FlickGameBack/pkg/engine/socket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	// corsの設定
	r.Use(cors.Default())

	// ルーティング
	r.GET("/word-get", core.WordGet())
	r.POST("/add-words", management.AddWord())
	r.POST("/register-score", score.RegisterScore())
	r.GET("/get-ranking", ranking.GetRanking())

	r.GET("/room-add", socket.AddRoomSession)
	r.GET("/room-join", socket.JoinRoomSession)
	// ソケット通信エンドポイント
	r.GET("/room/:id", socket.HandleWebSocket)
	r.GET("/room-result/:id", socket.HandleResultWebSocket)
	return r
}
