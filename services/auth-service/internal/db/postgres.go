package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/satoshi1975/smartChat/services/auth-service/config"
)

func NewPostgresDB(cfg *config.DBConfig) (*sql.DB, error) {
	dsn := "host=" + cfg.Host +
		" port=" + cfg.Port +
		" user=" + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.DBName +
		" sslmode=" + cfg.SSLMode

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
