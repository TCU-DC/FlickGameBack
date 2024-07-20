package model

type User struct {
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	NickName      string `json:"nick_name"`
	Password_hash string `json:"password_hash"`
}
