package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoshioka0101/go_auth/internal/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 汎用ハンドラのエンドポイント
	r.GET("/", handler.HelloWorldHandler)

	// 認証関連のエンドポイント
	 r.GET("/auth/google", handler.GoogleAuthHandler)
	 r.GET("/auth/google/callback", handler.GetAuthCallbackHandler)
	 r.POST("/auth/logout", handler.LogoutHandler)

	return r
}
