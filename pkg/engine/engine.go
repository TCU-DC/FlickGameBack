package engine

import (
	"FlickGameBack/pkg/engine/core"
	"FlickGameBack/pkg/engine/management"
	"FlickGameBack/pkg/engine/ranking"
	"FlickGameBack/pkg/engine/score"

	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {

	// ルーティング
	r.GET("/word-get", core.WordGet())
	r.POST("/add-words", management.AddWord())
	r.POST("/register-score", score.RegisterScore())
	r.GET("/get-ranking", ranking.GetRanking())
	return r
}
