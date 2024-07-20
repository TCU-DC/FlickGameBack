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
		`create table IF NOT EXISTS "level" (level_name text PRIMARY KEY, limit_time int)`,
		`create table IF NOT EXISTS "score" (score_id serial PRIMARY KEY, user_id text, point int, level text,FOREIGN KEY (user_id) REFERENCES "user"(user_id), FOREIGN KEY (level) REFERENCES "level"(level_name))`,
		`create table IF NOT EXISTS "word" (word_id text PRIMARY KEY, word_text text UNIQUE NOT NULL, word_furigana text NOT NULL, word_level text, point_allocation int NOT NULL, FOREIGN KEY (word_level) REFERENCES level(level_name))`,
	}

	for _, stm := range sql_stm {
		_, err := db.Exec(stm)
		if err != nil {
			slog.Error("failed to create table", "err=", err)
		}
	}
	// レベルテーブルのイニシャライズ
	InitLevelTable()
	slog.Info("create default table")
}

// レベルテーブルのイニシャライズ
func InitLevelTable() {
	db := db.Connect()
	defer db.Close()

	// レベルテーブルを削除し作成しなおす
	_, err := db.Exec(`delete from "level"`)
	if err != nil {
		slog.Error("failed to delete level", "err=", err)
	}

	// レベル名, 制限時間
	var sql_stm []string = []string{
		`insert into "level" (level_name, limit_time) values ('easy', 60)`,
		`insert into "level" (level_name, limit_time) values ('normal', 90)`,
		`insert into "level" (level_name, limit_time) values ('hard', 120)`,
	}

	for _, stm := range sql_stm {
		_, err := db.Exec(stm)
		if err != nil {
			slog.Error("failed to insert level", "err=", err)
		}
	}
	slog.Info("init level table")
}
