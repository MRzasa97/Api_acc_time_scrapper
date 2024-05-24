package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MRzasa97/Api_acc_time_scrapper/internal/tools"
)

type Env struct {
	db tools.DatabaseInterface
}

func (env *Env) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var bestTime tools.BestTime
	if err := json.NewDecoder(r.Body).Decode(&bestTime); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadGateway)
		return
	}

	if err := env.db.Create(bestTime); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Record saved!"})
}

func (env *Env) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	records, err := env.db.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(records)
}

func (env *Env) InitEnv(db tools.DatabaseInterface) *Env {
	env.db = db
	return env
}
