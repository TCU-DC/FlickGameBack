package read

import (
	"FlickGameBack/pkg/db"
	"log/slog"
)

// ユーザが存在するか確認する関数
func CheckExistUser(userID string) (bool, error) {
	db := db.Connect()
	defer db.Close()

	var count int
	err := db.QueryRow(`select count(*) from "user" where user_id = $1`, userID).Scan(&count)
	if err != nil {
		slog.Error("failed to check exist user", "err=", err)
		return false, err
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}
