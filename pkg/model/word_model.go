package model

type Word struct {
	WordId          string `json:"word_id"`
	WordText        string `json:"word_text"`
	WordLevel       string `json:"word_level"`
	PointAllocation int    `json:"point_allocation"`
}
