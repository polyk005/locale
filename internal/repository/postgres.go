package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
