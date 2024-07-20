package read

import (
	"FlickGameBack/pkg/db"
	"FlickGameBack/pkg/model"
)

// ランキングを取得
// レベル指定, 上位何位まで取得するか
func ReadRanking(level string, high_order int) ([]model.Ranking, error) {

	// DB接続
	db := db.Connect()
	defer db.Close()

	// ランキングを取得
	rows, err := db.Query(`SELECT user_id, point, level FROM "score" WHERE "level" = $1 ORDER BY point DESC LIMIT $2`, level, high_order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// ランキングを格納
	ranking := make([]model.Ranking, 0)
	for rows.Next() {
		var r model.Ranking
		err := rows.Scan(&r.UserID, &r.Score, &r.Level)
		if err != nil {
			return nil, err
		}
		ranking = append(ranking, r)
	}

	return ranking, nil
}
