package handler

import (
	"context"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/gin-gonic/gin"
)

// HelloWorldHandler は "Hello World" を返すハンドラ
func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// Google OAuth 認証開始のハンドラー
func GoogleAuthHandler(c *gin.Context) {
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", "google"))
	gothic.BeginAuthHandler(c.Writer, c.Request)
	// c.JSON(http.StatusOK, gin.H{"message": "auth/google成功したぜ"})
}

// Google OAuth コールバックのハンドラー
func GetAuthCallbackHandler(c *gin.Context) {
	// プロバイダーがクエリパラメータに含まれていない場合、"google" を追加する
	if c.Query("provider") == "" {
		q := c.Request.URL.Query()
		q.Set("provider", "google")
		c.Request.URL.RawQuery = q.Encode()
	}

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 認証成功後、フロントエンドのダッシュボードへリダイレクト
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/dashboard")
}

// LogoutHandler - セッション削除のハンドラー
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
