package model

type RoomJwt struct {
	RoomID      int
	LeaderName  string
	MemberNames []string
}

type RoomMessage struct {
	Nickname   string `json:"nickname"`
	MemberType string `json:"member_type"`
}
