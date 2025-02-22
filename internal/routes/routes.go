package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/yoshioka0101/go_auth/internal/auth"
	"github.com/yoshioka0101/go_auth/internal/handler"
)

// NewRouter はルートを登録した Gin エンジンを返します
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 汎用ハンドラのエンドポイント
	r.GET("/", handler.HelloWorldHandler)
	// r.GET("/health", handler.HealthHandler)

	// // 認証関連のエンドポイント
	// r.GET("/auth/google", auth.GoogleAuthHandler)
	// r.GET("/auth/google/callback", auth.CallbackHandler)
	// r.POST("/auth/logout", auth.LogoutHandler)

	return r
}
