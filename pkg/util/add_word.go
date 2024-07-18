package util

import (
	"FlickGameBack/pkg/db/create"
	"FlickGameBack/pkg/model"
)

// 単語を追加する
func AddWord() error {
	words := []model.Word{
		{WordId: "1", WordText: "猫", WordFurigana: "ネコ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "2", WordText: "犬", WordFurigana: "イヌ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "3", WordText: "鳥", WordFurigana: "トリ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "4", WordText: "魚", WordFurigana: "サカナ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "5", WordText: "熊", WordFurigana: "クマ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "6", WordText: "猿", WordFurigana: "サル", WordLevel: "normal", PointAllocation: 1},
		{WordId: "7", WordText: "狼", WordFurigana: "オオカミ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "8", WordText: "虎", WordFurigana: "トラ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "9", WordText: "獅子", WordFurigana: "シシ", WordLevel: "normal", PointAllocation: 1},
		{WordId: "10", WordText: "鹿", WordFurigana: "シカ", WordLevel: "normal", PointAllocation: 1},
	}

	for _, word := range words {
		if err := create.InsertWord(word); err != nil {
			return err
		}
	}
	return nil
}
