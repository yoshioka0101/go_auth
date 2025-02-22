package main

import (
	"log"

	"github.com/yoshioka0101/go_auth/internal/auth"
	"github.com/yoshioka0101/go_auth/internal/routes"
)

func main() {
	// 認証の初期化
	auth.NewAuth()

	// ルーターを生成する
	r := routes.NewRouter()

	// サーバーの起動
	log.Println("Starting server on :4000")
	if err := r.Run(":4000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
