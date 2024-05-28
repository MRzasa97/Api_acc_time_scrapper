package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MRzasa97/Api_acc_time_scrapper/internal/tools"
	"github.com/golang-jwt/jwt/v4"
)

type Env struct {
	db     tools.DatabaseInterface
	userDB tools.UserDatabaseInterface
	jwt    []byte
}

func (env *Env) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var bestTime tools.BestTime
	var track tools.Track
	if err := json.NewDecoder(r.Body).Decode(&bestTime); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadGateway)
		return
	}

	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID, err := extractUserNameFromToken(token.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bestTime.UserID = userID

	track, err = env.db.GetTrack(bestTime.TrackName)
	if err != nil {

		if err.Error() == "track not found" {

			track, err = env.db.CreateTrack(bestTime.TrackName)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}
		} else {

			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}

	if err := env.db.Create(bestTime, track.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "record created!"})
}

func (env *Env) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	records, err := env.db.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(records)
}

func (env *Env) InitEnv(db tools.DatabaseInterface, userDB tools.UserDatabaseInterface, jwt []byte) *Env {
	env.db = db
	env.userDB = userDB
	env.jwt = jwt
	return env
}

func extractUserNameFromToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		jwtSecret := getJwtKey()
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", string(tokenString))
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, fmt.Errorf("something went wrong with token parsing: %s", err)
	}
}

func getJwtKey() []byte {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("No JWT_SECRET_KEY variable in env file")
	}
	return []byte(jwtSecret)
}
