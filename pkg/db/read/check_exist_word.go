package read

import (
	"FlickGameBack/pkg/db"
)

// 単語が存在るかチェック
// 存在する場合はword_idを返す string
func CheckExistWord(word_text string) (string, error) {
	db := db.Connect()
	defer db.Close()

	var word_id string
	// word_idを取得
	err := db.QueryRow("SELECT word_id FROM word WHERE word_text=$1", word_text).Scan(&word_id)
	if err != nil {
		return "", err
	}

	return word_id, nil
}
