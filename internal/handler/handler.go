package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldHandler は "Hello World" を返すハンドラ
func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
