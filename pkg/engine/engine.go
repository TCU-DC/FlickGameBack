package engine

import (
	"FlickGameBack/pkg/engine/core"

	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	// ログとリカバリーのミドルウェアを設定(デフォルト)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// ルーティング
	r.GET("/word-get", core.WordGet())
	return r
}
