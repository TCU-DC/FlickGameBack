package update

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/model"
	"log/slog"
)

// 単語を更新する
func UpdateWord(word model.Word) error {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec("UPDATE word SET word_text=$1, word_furigana=$2, word_level=$3, point_allocation=$4 WHERE word_id=$5", word.WordText, word.WordFurigana, word.WordLevel, word.PointAllocation, word.WordId)
	if err != nil {
		slog.Error("failed to update word", "err=", err)
		return err
	}

	return nil
}
