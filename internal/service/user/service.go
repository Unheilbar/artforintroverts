package user

import (
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

type Storage interface {
	GetAll() ([]entities.User, error)
	Delete(ID string) error
	Update(ID string, user entities.User) error
}

type Cache interface {
	IsValid() bool
	SetInvalid()
	SetValid()
	GetAll() []entities.User
}

type StorageProvider interface {
	Storage() Storage
	Cache() Cache
}

type Service struct {
	StorageProvider
}

func NewService(sp StorageProvider) *Service {
	return &Service{
		StorageProvider: sp,
	}
}

func (s *Service) GetAll() ([]entities.User, error) {
	cache := s.StorageProvider.Cache()

	if cache.IsValid() {
		return cache.GetAll(), nil
	}

	return s.StorageProvider.Storage().GetAll()
}

func (s *Service) Delete(ID string) error {
	s.Cache().SetInvalid()
	err := s.Storage().Delete(ID)
	if err != nil {
		s.Cache().SetValid()
		return err
	}

	return nil
}

func (s *Service) Update(ID string, user entities.User) error {
	s.Cache().SetInvalid()
	err := s.Storage().Update(ID, user)
	if err != nil {
		s.Cache().SetValid()
		return err
	}

	return nil
}
