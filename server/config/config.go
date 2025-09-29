package config

import (
	"os"
)

var JwtSecret []byte

func InitConfig() {
	// беремо з ENV, якщо нема — fallback
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretjwtkey" // fallback
	}
	JwtSecret = []byte(secret)
}
