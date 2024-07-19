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
		`create table IF NOT EXISTS "level" (level_name text PRIMARY KEY)`,
		`create table IF NOT EXISTS "score" (score_id serial PRIMARY KEY, user_id text, point int, average_speed float, level text, FOREIGN KEY (level) REFERENCES level(level_name))`,
		`create table IF NOT EXISTS "word" (word_id text PRIMARY KEY, word_text text UNIQUE NOT NULL, word_furigana text NOT NULL, word_level text, point_allocation int NOT NULL, FOREIGN KEY (word_level) REFERENCES level(level_name))`,
	}

	// レベルのデフォルト値を挿入easy, normal, hard
	_, err := db.Exec(`insert into level (level_name) values ('easy'), ('normal'), ('hard')`)
	if err != nil {
		slog.Error("failed to insert default level", "err=", err)
	}

	for _, stm := range sql_stm {
		_, err := db.Exec(stm)
		if err != nil {
			slog.Error("failed to create table", "err=", err)
		}
	}
	slog.Info("create default table")
}
