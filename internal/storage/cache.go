package storage

import (
	"sync"

	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

type cache struct {
	isValid bool
	mx      *sync.Mutex
	users   []entities.User
}

func (c *cache) IsValid() bool {
	c.lock()
	defer c.unlock()
	return c.isValid
}

func (c *cache) GetAll() []entities.User {
	return c.users
}

func (c *cache) SetInvalid() {
	c.lock()
	c.isValid = false
	c.unlock()
}

func (c *cache) SetValid() {
	c.lock()
	c.isValid = true
	c.unlock()
}

// cache should be locked outside this method
func (c *cache) refresh(users []entities.User) {
	c.users = users
	c.isValid = true
}

func (c *cache) lock() {
	c.mx.Lock()
}

func (c *cache) unlock() {
	c.mx.Unlock()
}
