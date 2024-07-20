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
		`create table IF NOT EXISTS "user" (user_id text PRIMARY KEY, username text UNIQUE NOT NULL, nickname text NOT NULL, password_hash text)`,
		`create table IF NOT EXISTS "level" (level_name text PRIMARY KEY)`,
		`create table IF NOT EXISTS "score" (score_id serial PRIMARY KEY, user_id text, point int, level text,FOREIGN KEY (user_id) REFERENCES "user"(user_id), FOREIGN KEY (level) REFERENCES "level"(level_name))`,
		`create table IF NOT EXISTS "word" (word_id text PRIMARY KEY, word_text text UNIQUE NOT NULL, word_furigana text NOT NULL, word_level text, point_allocation int NOT NULL, FOREIGN KEY (word_level) REFERENCES level(level_name))`,
	}

	// 存在しなければ, レベルのデフォルト値を挿入easy, normal, hard
	_, err := db.Exec(`insert into level (level_name) select 'easy' where not exists (select * from level where level_name = 'easy')`)
	if err != nil {
		slog.Error("failed to insert level", "err=", err)
	}
	_, err = db.Exec(`insert into level (level_name) select 'normal' where not exists (select * from level where level_name = 'normal')`)
	if err != nil {
		slog.Error("failed to insert level", "err=", err)
	}
	_, err = db.Exec(`insert into level (level_name) select 'hard' where not exists (select * from level where level_name = 'hard')`)
	if err != nil {
		slog.Error("failed to insert level", "err=", err)
	}

	for _, stm := range sql_stm {
		_, err := db.Exec(stm)
		if err != nil {
			slog.Error("failed to create table", "err=", err)
		}
	}
	slog.Info("create default table")
}
