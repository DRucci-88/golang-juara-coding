package service

import (
	"materi-middleware-gorm/models"
	"materi-middleware-gorm/repository"
)

type ConcertService interface {
	CreateConcert(concert *models.Concert) error
	GetAllConcerts() ([]models.Concert, error)
	GetConcertByID(id int) (models.Concert, error)
	UpdateConcert(concert *models.Concert) error
	DeleteConcert(id int) error
}

type concertService struct {
	repo repository.ConcertRepository
}

func NewConcertService(repo repository.ConcertRepository) ConcertService {
	return &concertService{repo: repo}
}

func (s *concertService) CreateConcert(concert *models.Concert) error {
	return s.repo.Create(concert)
}

func (s *concertService) GetAllConcerts() ([]models.Concert, error) {
	return s.repo.FindAll()
}

func (s *concertService) GetConcertByID(id int) (models.Concert, error) {
	return s.repo.FindByID(id)
}

func (s *concertService) UpdateConcert(concert *models.Concert) error {
	return s.repo.Update(concert)
}

func (s *concertService) DeleteConcert(id int) error {
	return s.repo.Delete(id)
}
