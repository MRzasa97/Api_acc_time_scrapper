package tools

import (
	"database/sql"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
)

type PostgresUserDB struct {
	DB *sql.DB
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewPostgresUserDB(connStr string) (*PostgresUserDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresUserDB{DB: db}, nil
}

func (pg *PostgresUserDB) CreateUser(user User) error {
	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	`
	err := pg.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	return nil
}

func (pg *PostgresUserDB) GetUser(username string) (User, error) {
	var user User

	query := `
		SELECT id, username, password
		FROM users
		WHERE username = $1
	`
	err := pg.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return user, nil
}

func (pg *PostgresUserDB) GetTrack(trackName string) (Track, error) {
	var track Track

	query := `
		SELECT id, track_name
		FROM tracks
		WHERE track_name = $1
	`

	err := pg.DB.QueryRow(query, trackName).Scan(&track.ID, &track.TrackName)
	if err != nil {
		if err == sql.ErrNoRows {
			return track, fmt.Errorf("track not found")
		}
		return track, fmt.Errorf("failed to retrieve track: %v", err)
	}

	return track, nil
}

func (pg *PostgresUserDB) CreateTrack(trackName string) (Track, error) {
	query := `
		INSERT INTO tracks (track_name)
		VALUES ($1)
		RETURNING id, track_name
	`
	var track Track
	err := pg.DB.QueryRow(query, trackName).Scan(&track.ID, &track.TrackName)
	if err != nil {
		return track, fmt.Errorf("failed to insert track record: %s", err)
	}
	return track, nil
}

func (pg *PostgresUserDB) Create(bt BestTime, trackID int) error {
	query := `
		INSERT INTO best_times (user_id, car, track_id, best_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var newID int
	err := pg.DB.QueryRow(query, bt.UserID, bt.CarModel, trackID, bt.BestTime).Scan(&newID)
	if err != nil {
		return fmt.Errorf("failed to insert record: %s", err)
	}

	return nil
}

func (pg *PostgresUserDB) GetAll() ([]BestTime, error) {
	query := `
		SELECT bt.car, bt.best_time, t.track_name
		FROM best_times bt
		JOIN
		tracks t on bt.track_id = t.id
	`
	rows, err := pg.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error quering the database: %w", err)
	}

	defer rows.Close()

	var bestTimes []BestTime

	for rows.Next() {
		var bestTime BestTime
		err := rows.Scan(&bestTime.CarModel, &bestTime.BestTime, &bestTime.TrackName)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		bestTimes = append(bestTimes, bestTime)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows %w", err)
	}

	return bestTimes, nil
}
