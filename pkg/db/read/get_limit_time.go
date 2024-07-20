package read

import (
	"FlickGameBack/pkg/db"
	"log/slog"
)

// 指定されたレベルの制限時間を取得する関数
func ReadLimitTime(level string) (int, error) {
	db := db.Connect()
	defer db.Close()

	var limitTime int
	err := db.QueryRow(`select limit_time from level where level_name = $1`, level).Scan(&limitTime)
	if err != nil {
		slog.Error("failed to read limit time", "err=", err)
		return 0, err
	}
	return limitTime, nil
}
