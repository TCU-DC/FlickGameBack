package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketをアップグレードするための設定
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	rooms    = make(map[string]map[*websocket.Conn]Message) // Connとニックネームのマップに変更
	roomLock = sync.Mutex{}
)

type Message struct {
	Room       string `json:"room"`
	Nickname   string `json:"nickname"`
	MemberType string `json:"member_type"`
	Action     string `json:"action"`
}

// 参加者リストを更新する関数
func broadcastParticipants(roomID string) {
	participants := make([]Message, 0)
	roomLock.Lock()
	for _, message := range rooms[roomID] {
		participants = append(participants, message)
	}
	roomLock.Unlock()

	message, _ := json.Marshal(participants)

	roomLock.Lock()
	for client := range rooms[roomID] {
		client.WriteMessage(websocket.TextMessage, message)
	}
	roomLock.Unlock()
}

// WebSocket接続を処理するハンドラ関数
func HandleWebSocket(c *gin.Context) {
	roomID := c.Param("id")
	slog.Info(fmt.Sprintf("Room ID: %s", roomID))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to upgrade to websocket: %v", err))
		c.String(http.StatusBadRequest, "Failed to upgrade to websocket: %v", err)
		return
	}
	defer conn.Close()
	slog.Info(fmt.Sprintf("Client connected: %v", conn.RemoteAddr()))

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Info(fmt.Sprintf("Client disconnected: %v", conn.RemoteAddr()))
			roomLock.Lock()
			delete(rooms[roomID], conn)
			if len(rooms[roomID]) == 0 {
				delete(rooms, roomID)
				slog.Info(fmt.Sprintf("Room deleted: %s", roomID))
			}
			roomLock.Unlock()
			broadcastParticipants(roomID)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			slog.Info(fmt.Sprintf("Failed to unmarshal message: %v", err))
			continue
		}

		if msg.Action == "join" {
			roomLock.Lock()
			if _, ok := rooms[roomID]; !ok {
				rooms[roomID] = make(map[*websocket.Conn]Message)
			}

			if len(rooms[roomID]) == 0 {
				msg.MemberType = "leader"
			} else {
				msg.MemberType = "member"
			}

			rooms[roomID][conn] = msg
			roomLock.Unlock()

			slog.Info(fmt.Sprintf("Client added to room: %s as %s", roomID, msg.Nickname))
			broadcastParticipants(roomID)
		} else if msg.Action == "start" {
			fmt.Println("Start message received from leader, broadcasting to all clients")
			roomLock.Lock()
			for client := range rooms[roomID] {
				err = client.WriteMessage(websocket.TextMessage, []byte("start"))
				if err != nil {
					slog.Info(fmt.Sprintf("Failed to send start message to client: %v", client.RemoteAddr()))
					client.Close()
					delete(rooms[roomID], client)
				}
			}
			roomLock.Unlock()
		} else {
			slog.Info(fmt.Sprintf("Received message from client: %s", message))

			roomLock.Lock()
			for client := range rooms[roomID] {
				if client != conn {
					err = client.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						slog.Info(fmt.Sprintf("Failed to send message to client: %v", client.RemoteAddr()))
						client.Close()
						delete(rooms[roomID], client)
					}
				}
			}
			roomLock.Unlock()
		}
	}
}
