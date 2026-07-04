package config

import (
	"log"
	"os"
)

// JWTKey returns the JWT signing key from the JWT_SECRET environment variable.
// It panics at startup if the variable is not set, to fail fast in misconfigured environments.
func JWTKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable is not set")
	}
	return []byte(secret)
}
