package socket

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketをアップグレードするための設定
var (
	res_upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	res_rooms    = make(map[string]map[*websocket.Conn]ResMessage) // Connとニックネームのマップに変更
	res_roomLock = sync.Mutex{}
)

type ResMessage struct {
	Room       string `json:"room"`
	Nickname   string `json:"nickname"`
	MemberType string `json:"member_type"`
	Action     string `json:"action"`
	Score      string `json:"score"`
}

// HandleResultWebSocket は結果を送信するためのWebSocket接続を処理するハンドラ関数です。
func HandleResultWebSocket(c *gin.Context) {
	roomID := c.Param("id")
	slog.Info(fmt.Sprintf("Room ID: %s", roomID))

	conn, err := res_upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to upgrade to websocket: %v", err))
		return
	}
	defer conn.Close() // 接続が閉じられることを保証

	res_roomLock.Lock()
	if _, ok := res_rooms[roomID]; !ok {
		res_rooms[roomID] = make(map[*websocket.Conn]ResMessage)
	}
	res_rooms[roomID][conn] = ResMessage{} // 初期値として空のメッセージを追加
	res_roomLock.Unlock()

	for {
		var message ResMessage
		err := conn.ReadJSON(&message)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to read message: %v", err))
			break
		}

		res_roomLock.Lock()
		res_rooms[roomID][conn] = message

		// 現在のルーム内のすべてのスコアを収集し、クライアントに送信
		for client := range res_rooms[roomID] {
			if err := client.WriteJSON(message); err != nil {
				slog.Error(fmt.Sprintf("Failed to write message to client: %v", err))
			}
		}
		res_roomLock.Unlock()
	}

	res_roomLock.Lock()
	delete(res_rooms[roomID], conn)
	res_roomLock.Unlock()
}
