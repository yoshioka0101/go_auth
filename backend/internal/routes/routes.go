package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/go_auth/backend/internal/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// CORS 設定を追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// 汎用ハンドラのエンドポイント
	r.GET("/", handler.HelloWorldHandler)

	// 認証関連のエンドポイント
	r.GET("/auth/google", handler.GoogleAuthHandler)
	r.GET("/auth/google/callback", handler.GetAuthCallbackHandler)
	r.GET("/auth/session", handler.SessionHandler)
	r.POST("/auth/logout", handler.LogoutHandler)

	return r
}
