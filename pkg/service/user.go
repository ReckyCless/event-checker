package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUserRole(userID int, role app.UserRoleInput) (int64, error) {
	return s.repo.UpdateUserRole(userID, role)
}

func (s *UserService) UpdateUserOrganisator(userID int, organisator app.UserOrganisatorInput) (int64, error) {
	return s.repo.UpdateUserOrganisator(userID, organisator)
}
