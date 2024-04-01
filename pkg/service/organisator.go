package service

import (
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/repository"
)

type OrganistorService struct {
	repo repository.Organisator
}

func NewOrganisatorService(repo repository.Organisator) *OrganistorService {
	return &OrganistorService{repo: repo}
}

func (s *OrganistorService) Create(organisator app.Organisator) (int, error) {
	return s.repo.Create(organisator)
}

func (s *OrganistorService) GetAll() ([]app.Organisator, error) {
	return s.repo.GetAll()
}

func (s *OrganistorService) GetByID(organisatorID int) (app.Organisator, error) {
	return s.repo.GetByID(organisatorID)
}

func (s *OrganistorService) Delete(organisatorID int) (int64, error) {
	return s.repo.Delete(organisatorID)
}

func (s *OrganistorService) Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error) {
	return s.repo.Update(organisatorID, input)
}
