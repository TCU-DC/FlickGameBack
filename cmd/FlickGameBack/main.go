package main

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/engine"
	"FlickGameBack/pkg/util"

	"github.com/gin-gonic/gin"
)

func init() {
	// デフォルトのテーブルを作成
	create.CreateDefaultTable()

	// デフォルトの単語を追加
	util.AddWord()
}

func main() {
	r := gin.Default()
	r = engine.Engine(r)
	r.Run(":8080")
}
