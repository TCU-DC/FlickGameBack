package socket

import (
	"FlickGameBack/pkg/util"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

const SECRET_KEY = "SECRET"

// ROOMセッションに情報を登録する
func AddRoomSession(c *gin.Context) {
	// リクエストボディの取得
	leaderName := c.Query("leader_name")
	if leaderName == "" {
		slog.Error("leader_name is empty")
		c.JSON(400, gin.H{"error": "leader_name is empty"})
	}

	// ルームIDの生成
	roomID := util.RandomInt(5)

	// セッション情報の登録
	// JWTの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"room_id":     roomID,
		"leaderName":  leaderName,
		"memberNames": []string{},
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
		return
	}

	// ヘッダにJWTを設定
	c.Header("Authorization", tokenString)

	// レスポンス
	c.JSON(200, gin.H{
		"room_id":  roomID,
		"nickname": leaderName,
		"message":  "you are leader",
	})
}

// ROOMセッションに参加する
func JoinRoomSession(c *gin.Context) {
	// リクエストボディの取得
	roomID := c.Query("room_id")
	myName := c.Query("my_name")
	if roomID == "" {
		slog.Error("room_id is empty")
		c.JSON(400, gin.H{"error": "room_id is empty"})
	}
	if myName == "" {
		slog.Error("my_name is empty")
		c.JSON(400, gin.H{"error": "my_name is empty"})
	}

	// ルームIDをintに変換
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		slog.Error("room_id is invalid")
		c.JSON(400, gin.H{"error": "room_id is invalid"})
	}

	// セッション情報の登録
	// JWTの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"room_id":     roomIDInt,
		"myname":      myName,
		"leaderName":  "",
		"memberNames": []string{myName},
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
		return
	}

	// ヘッダにJWTを設定
	c.Header("Authorization", tokenString)

	// レスポンス
	c.JSON(200, gin.H{
		"room_id":  roomID,
		"nickname": myName,
		"message":  "you are member",
	})
}
