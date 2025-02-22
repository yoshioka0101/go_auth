package handler

import (
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/gin-gonic/gin"
)

// HelloWorldHandler は "Hello World" を返すハンドラ
func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// GoogleOAuth認証開始のハンドラーを追加
func GoogleAuthHandler(c *gin.Context) {
	c.String(http.StatusOK, "GoogleOAuthHandler")
}

// Logoutのハンドラーを追加
func LogoutHandler(c *gin.Context) {
    // セッション削除などの実装
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

// GoogleOAuthにCallbackするためのハンドラーを追加
func GetAuthCallbackHandler(c *gin.Context) {
    user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user": user})
}
