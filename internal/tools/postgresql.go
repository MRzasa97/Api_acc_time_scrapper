package tools

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresUserDB struct {
	DB *sql.DB
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
