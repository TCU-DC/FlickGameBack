package main

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/engine"

	"github.com/gin-gonic/gin"
)

func init() {
	// デフォルトのテーブルを作成
	create.CreateDefaultTable()
}

func main() {
	r := gin.Default()
	r = engine.Engine(r)
	r.Run(":8080")
}
