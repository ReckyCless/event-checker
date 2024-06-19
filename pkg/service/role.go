package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type RoleService struct {
	repo repository.Role
}

func NewRoleService(repo repository.Role) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetAll() ([]app.Role, error) {
	return s.repo.GetAll()
}

func (s *RoleService) GetByID(roleID int) (app.Role, error) {
	return s.repo.GetByID(roleID)
}

func (s *RoleService) Update(roleID int, input app.UpdateRoleInput) (int64, error) {
	return s.repo.Update(roleID, input)
}
