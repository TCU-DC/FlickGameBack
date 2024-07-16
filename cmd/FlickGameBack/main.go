package main

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/engine"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのテーブルを作成
	create.CreateDefaultTable()
	r := gin.Default()
	r = engine.Engine(r)
	r.Run(":8080")
}
