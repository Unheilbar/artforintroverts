package adapter

import (
	"github.com/unheilbar/artforintrovert_entry_task/internal/service/user"
	"github.com/unheilbar/artforintrovert_entry_task/internal/storage"
)

type UserStorageProviderAdapter struct {
	provider *storage.Provider
}

func (u *UserStorageProviderAdapter) Storage() user.Storage {
	return u.provider.Storage()
}

func (u *UserStorageProviderAdapter) Cache() user.Cache {
	return u.provider.Cache()
}

func NewStorageProviderAdapter(provider *storage.Provider) *UserStorageProviderAdapter {
	return &UserStorageProviderAdapter{
		provider: provider,
	}
}
