package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type OrganisatorService struct {
	repo repository.Organisator
}

func NewOrganisatorService(repo repository.Organisator) *OrganisatorService {
	return &OrganisatorService{repo: repo}
}

func (s *OrganisatorService) Create(organisator app.CreateOrganisatorInput) (int, error) {
	return s.repo.Create(organisator)
}

func (s *OrganisatorService) GetAll() ([]app.GetOrganisatorOutput, error) {
	return s.repo.GetAll()
}

func (s *OrganisatorService) GetByID(organisatorID int) (app.GetOrganisatorOutput, error) {
	return s.repo.GetByID(organisatorID)
}

func (s *OrganisatorService) Delete(organisatorID int) (int64, error) {
	return s.repo.Delete(organisatorID)
}

func (s *OrganisatorService) Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error) {
	return s.repo.Update(organisatorID, input)
}
