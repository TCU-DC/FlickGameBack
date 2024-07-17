package model

type Word struct {
	WordId          string `json:"word_id"`
	WordText        string `json:"word_text"`
	WordFurigana    string `json:"word_furigana"`
	WordLevel       string `json:"word_level"`
	PointAllocation int    `json:"point_allocation"`
}
