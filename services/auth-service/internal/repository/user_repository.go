package repository

import (
	"context"
	"database/sql"
	"github.com/satoshi1975/smartChat/services/auth-service/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE username = $1"
	row := r.db.QueryRowContext(ctx, query, username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
