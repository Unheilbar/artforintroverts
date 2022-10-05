package user

import (
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

type Storage interface {
	GetAll() ([]entities.User, error)
	Delete(ID string) error
	UpdateAndReturn(ID string, user entities.User) (entities.User, error)
	GetById(ID string) (entities.User, error)
	Save(entities.User) error
}

type Cache interface {
	Set(user entities.User)
	Get(key string) (entities.User, bool)
	Delete(key string)
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
	return s.StorageProvider.Storage().GetAll()
}

func (s *Service) GetById(ID string) (entities.User, error) {
	user, ok := s.Cache().Get(ID)
	if ok {
		return user, nil
	}

	user, err := s.Storage().GetById(ID)
	if err != nil {
		return entities.User{}, err
	}

	s.Cache().Set(user)

	return user, nil
}

func (s *Service) Delete(ID string) error {
	err := s.Storage().Delete(ID)
	if err != nil {
		return err
	}

	s.Cache().Delete(ID)
	return nil
}

func (s *Service) Update(ID string, user entities.User) error {
	updated, err := s.Storage().UpdateAndReturn(ID, user)
	if err != nil {
		return err
	}

	s.Cache().Set(updated)
	return nil
}

func (s *Service) Save(user entities.User) error {
	err := s.Storage().Save(user)
	if err != nil {
		return err
	}

	s.Cache().Set(user)
	return nil
}
