package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MRzasa97/Api_acc_time_scrapper/internal/handlers"
	"github.com/MRzasa97/Api_acc_time_scrapper/internal/tools"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Text string
}

func getJwtKey() []byte {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("No JWT_SECRET_KEY variable in env file")
	}
	return []byte(jwtSecret)
}

func main() {
	dbUser, err := tools.NewPostgresUserDB("postgresql://myuser:mypassword@db/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	db := dbUser
	env := &handlers.Env{}
	jwtSecret := getJwtKey()
	env.InitEnv(db, dbUser, jwtSecret)

	var router *chi.Mux = chi.NewRouter()
	fmt.Println("Starting GO API service...")
	handlers.Handler(router, env)

	err = http.ListenAndServe(":8000", router)

	if err != nil {
		log.Error("Error!")
	}
}
