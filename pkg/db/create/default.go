package create

import (
	"FlickGameBack/pkg/db"
	"log/slog"

	_ "github.com/lib/pq"
)

// デフォルトのテーブルを作成する
func CreateDefaultTable() {
	db := db.Connect()
	defer db.Close()

	var sql_stm []string = []string{
		`create table IF NOT EXISTS "word" (word_id text PRIMARY KEY, word_text text, word_furigana text, word_level text, point_allocation int)`,
	}

	for _, stm := range sql_stm {
		_, err := db.Exec(stm)
		if err != nil {
			slog.Error("failed to create table", "err=", err)
		}
	}
	slog.Info("create default table")
}
