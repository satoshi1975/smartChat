package services

import (
	"context"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/models"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/repository"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile *models.Profile) error {
	return s.repo.CreateProfile(ctx, profile)
}

func (s *ProfileService) GetProfileByID(ctx context.Context, id int) (*models.Profile, error) {
	return s.repo.GetProfileByID(ctx, id)
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	return s.repo.UpdateProfile(ctx, profile)
}

func (s *ProfileService) DeleteProfile(ctx context.Context, id int) error {
	return s.repo.DeleteProfile(ctx, id)
}

func (s *ProfileService) AddFriend(ctx context.Context, profileID, friendID int) error {
	return s.repo.AddFriend(ctx, profileID, friendID)
}

func (s *ProfileService) BlockUser(ctx context.Context, profileID, blockedID int) error {
	return s.repo.BlockUser(ctx, profileID, blockedID)
}
