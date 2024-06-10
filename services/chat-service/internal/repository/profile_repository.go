package repository

import (
	"context"
	"database/sql"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/models"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *models.Profile) error {
	query := "INSERT INTO profiles (user_id, first_name, last_name, bio) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.db.QueryRowContext(ctx, query, profile.UserID, profile.FirstName, profile.LastName, profile.Bio).Scan(&profile.ID)
}

func (r *ProfileRepository) GetProfileByID(ctx context.Context, id int) (*models.Profile, error) {
	query := "SELECT id, user_id, first_name, last_name, bio FROM profiles WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var profile models.Profile
	if err := row.Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.LastName, &profile.Bio); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	query := "UPDATE profiles SET first_name = $2, last_name = $3, bio = $4 WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, profile.ID, profile.FirstName, profile.LastName, profile.Bio)
	return err
}

func (r *ProfileRepository) DeleteProfile(ctx context.Context, id int) error {
	query := "DELETE FROM profiles WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ProfileRepository) AddFriend(ctx context.Context, profileID, friendID int) error {
	query := "INSERT INTO friends (profile_id, friend_id) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, profileID, friendID)
	return err
}

func (r *ProfileRepository) BlockUser(ctx context.Context, profileID, blockedID int) error {
	query := "INSERT INTO blocked (profile_id, blocked_id) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, profileID, blockedID)
	return err
}
