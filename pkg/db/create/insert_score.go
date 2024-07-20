package create

import (
	"FlickGameBack/pkg/db"
	"log/slog"
)

// 得点情報を登録する関数
func InsertScore(userID string, point int, level string) error {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(`insert into score (user_id, point, level) values ($1, $2, $3)`, userID, point, level)
	if err != nil {
		slog.Error("failed to insert score", "err=", err)
		return err
	}
	return nil
}
