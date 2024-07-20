package model

type Ranking struct {
	UserID   string `json:"user_id"`
	NickName string `json:"nickname"`
	Score    int    `json:"score"`
}
