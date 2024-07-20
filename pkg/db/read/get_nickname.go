package read

import (
	"FlickGameBack/pkg/db"
)

// ユーザーIDからニックネームを取得
func GetNickname(userID string) (string, error) {
	// DB接続
	db := db.Connect()
	defer db.Close()

	// ニックネームを取得
	var nickname string
	err := db.QueryRow(`SELECT nickname FROM "user" WHERE "user_id" = $1`, userID).Scan(&nickname)
	if err != nil {
		return "", err
	}

	return nickname, nil
}
