package tools

import (
	"log"
	"os"
)

func getJwtKey() []byte {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("No JWT_SECRET_KEY variable in env file")
	}
	return []byte(jwtSecret)
}
