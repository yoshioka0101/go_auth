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
}

// Google OAuth コールバックのハンドラー
func GetAuthCallbackHandler(c *gin.Context) {
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

	// セッションにユーザー情報を保存
	session, _ := gothic.Store.Get(c.Request, "_gothic_session")
	session.Values["user"] = gin.H{
		"name":   user.Name,
		"email":  user.Email,
		"avatar": user.AvatarURL,
	}
	session.Save(c.Request, c.Writer)

	// 認証成功メッセージをJSONで返す
	c.JSON(http.StatusOK, gin.H{"message": "login成功しました！"})
}

// SessionHandler - セッション情報を取得する
func SessionHandler(c *gin.Context) {
	session, _ := gothic.Store.Get(c.Request, "_gothic_session")
	user, exists := session.Values["user"]

	if !exists {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user":          user,
	})
}

// LogoutHandler - セッション削除のハンドラー
func LogoutHandler(c *gin.Context) {
    // auth_token Cookieを削除
    c.SetCookie("auth_token", "", -1, "/", "localhost", false, true)

    // ログアウト成功のメッセージを返す
    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
