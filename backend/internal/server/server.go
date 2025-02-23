package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

// Server はサーバーの状態を管理するための構造体です。
// 今回はフィールドは必要なければ空のままでOKです。
type Server struct{}

// NewServer は新しい Server インスタンスを返します。
func NewServer() *Server {
	return &Server{}
}

// HelloWorldHandler は Server のメソッドとして実装した例です。
func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := map[string]string{"message": "Hello World from Server"}
	c.JSON(http.StatusOK, resp)
}

// GoogleAuthHandler は Server のメソッドとして Google OAuth 認証開始処理を実装した例です。
func (s *Server) GoogleAuthHandler(c *gin.Context) {
	// provider を "google" に設定する
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", "google"))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
