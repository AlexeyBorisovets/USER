package service

import (
	"github.com/AlexeyBorisovets/USER/internal/model"
	"github.com/AlexeyBorisovets/USER/internal/repository"

	"context"
)

// Service s
type Service struct {
	jwtKey []byte
	repos  repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository, jwtKey []byte) *Service {
	return &Service{repos: pool, jwtKey: jwtKey}
}

// GetUser _
func (se *Service) GetUser(ctx context.Context, id string) (*model.User, error) {
	return se.repos.GetUserByID(ctx, id)
}

// GetAllUsers _
func (se *Service) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return se.repos.GetAllUsers(ctx)
}

// DeleteUser _
func (se *Service) DeleteUser(ctx context.Context, id string) error {
	return se.repos.DeleteUser(ctx, id)
}

// UpdateUser _
func (se *Service) UpdateUser(ctx context.Context, id string, user *model.User) error {
	return se.repos.UpdateUser(ctx, id, user)
}

// GetUserByUserType _
func (se *Service) GetUserByUserType(ctx context.Context, usertype string) ([]*model.User, error) {
	return se.repos.GetUserByUserType(ctx, usertype)
}

// GetBalanceByID _
func (se *Service) GetBalanceByID(ctx context.Context, userId string) (uint, error) {
	return se.repos.GetBalanceByID(ctx, userId)
}

// UpdateBalance _
func (se *Service) UpdateBalance(ctx context.Context, userId string, balance uint) error {
	return se.repos.UpdateBalance(ctx, userId, balance)
}
