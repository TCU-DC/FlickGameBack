package create

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/model"
	"log/slog"
)

// 単語を追加する
func InsertWord(word model.Word) error {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO word (word_id, word_text, word_furigana, word_level, point_allocation) VALUES ($1, $2, $3, $4, $5)", word.WordId, word.WordText, word.WordFurigana, word.WordLevel, word.PointAllocation)
	if err != nil {
		slog.Error("failed to insert word", "err=", err)
		return err
	}

	return nil
}
