package read

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/model"
	"log/slog"
)

// 単語を取得する
func ReadWords(level string, count int) []model.Word {
	db := db.Connect()
	defer db.Close()

	var words []model.Word
	rows, err := db.Query("SELECT * FROM word WHERE word_level = $1 ORDER BY RANDOM() LIMIT $2", level, count)
	if err != nil {
		slog.Error("failed to select words", "err=", err)
	}

	for rows.Next() {
		var word model.Word
		err := rows.Scan(&word.WordId, &word.WordText, &word.WordLevel, &word.PointAllocation)
		if err != nil {
			slog.Error("failed to scan word", "err=", err)
		}
		words = append(words, word)
	}

	return words
}
