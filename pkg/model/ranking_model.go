package model

type Ranking struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
	Level    string `json:"level"`
}
