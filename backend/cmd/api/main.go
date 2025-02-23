package main

import (
	"log"

	"github.com/yoshioka0101/go_auth/backend/internal/auth"
	"github.com/yoshioka0101/go_auth/backend/internal/routes"
)

func main() {
	// 認証の初期化
	auth.NewAuth()

	r := routes.NewRouter()

	// サーバーの起動
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
