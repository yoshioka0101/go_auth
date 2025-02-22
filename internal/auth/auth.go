package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 84600 * 30
	IsProd = false
)

func NewAuth() {
	// .env の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URI")

	// Google OAuth プロバイダの初期化
	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, googleRedirectURL, "email", "profile"),
	)
}
