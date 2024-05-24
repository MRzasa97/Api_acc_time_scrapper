package main

import (
	"fmt"
	"net/http"

	"github.com/MRzasa97/Api_acc_time_scrapper/internal/handlers"
	"github.com/MRzasa97/Api_acc_time_scrapper/internal/tools"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Text string
}

func main() {
	db := tools.InitMockDB()
	env := &handlers.Env{}
	env.InitEnv(db)

	var router *chi.Mux = chi.NewRouter()
	fmt.Println("Starting GO API service...")
	handlers.Handler(router, env)

	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Error("Error!")
	}
}
