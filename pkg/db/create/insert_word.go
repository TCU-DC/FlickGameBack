package create

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/db/read"
	"FlickGameBack/pkg/db/update"
	"FlickGameBack/pkg/model"
	"log/slog"
)

// 単語を追加する
func InsertWord(word model.Word) error {
	db := db.Connect()
	defer db.Close()

	// 既に存在する単語であれば更新する
	if word_id, err := read.CheckExistWord(word.WordText); err == nil && word_id == word.WordId {
		if err := update.UpdateWord(word); err != nil {
			return err
		}
		return nil
	}

	_, err := db.Exec("INSERT INTO word (word_id, word_text, word_furigana, word_level, point_allocation) VALUES ($1, $2, $3, $4, $5)", word.WordId, word.WordText, word.WordFurigana, word.WordLevel, word.PointAllocation)
	if err != nil {
		slog.Error("failed to insert word", "err=", err)
		return err
	}

	return nil
}
