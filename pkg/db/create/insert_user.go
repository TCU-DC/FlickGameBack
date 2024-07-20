package create

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/db/read"
	"FlickGameBack/pkg/model"
	"log/slog"
)

// ユーザ情報を登録する関数
func InsertUser(user model.User) error {
	db := db.Connect()
	defer db.Close()

	// ユーザが存在しない場合のみ登録
	exist, err := read.CheckExistUser(user.UserID)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	_, err = db.Exec(`insert into "user" (user_id, username, nickname, password_hash) values ($1, $2, $3, $4)`, user.UserID, user.UserName, user.NickName, user.Password_hash)
	if err != nil {
		slog.Error("failed to insert user", "err=", err)
		return err
	}
	return nil
}
