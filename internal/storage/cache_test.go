package storage

import (
	"fmt"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

const cap = 4

func Test__cache(t *testing.T) {
	cache := prepareSomeCache()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tst := assert.New(t)

	t.Run("save with no prune", func(t *testing.T) {
		cache.Set(entities.User{
			ID:       "3",
			Nickname: "crosh3",
			Age:      3,
		})
		tst.Equal(cache.queue.size, uint(4))
	})

	t.Run("save with prune", func(t *testing.T) {
		expTail := entities.User{
			ID:       "4",
			Nickname: "crosh4",
			Age:      4,
		}

		expHead := entities.User{
			ID:       "1",
			Nickname: "crosh1",
			Age:      1,
		}
		cache.Set(expTail)

		tst.Equal(cache.queue.size, uint(4))
		tst.Equal(cache.queue.tail.val, expTail)
		tst.Equal(cache.queue.head.val, expHead)
	})

	t.Run("set existed user", func(t *testing.T) {
		expTail := entities.User{
			ID:       "2",
			Nickname: "crosh200",
			Age:      200,
		}

		cache.Set(expTail)

		tst.Equal(cache.queue.size, uint(4))
		tst.Equal(cache.queue.tail.val, expTail)
	})

	t.Run("rm head user", func(t *testing.T) {
		headId := cache.queue.head.key
		expHead := cache.queue.head.next

		cache.Delete(headId)

		tst.Equal(cache.queue.size, uint(3))
		tst.Equal(cache.queue.head, expHead)
	})

	t.Run("rm tail user", func(t *testing.T) {
		tailId := cache.queue.tail.key
		expTail := cache.queue.tail.prev

		cache.Delete(tailId)

		tst.Equal(cache.queue.size, uint(2))
		tst.Equal(cache.queue.tail, expTail)
	})

	t.Run("get head user", func(t *testing.T) {
		headId := cache.queue.head.key
		expUser := cache.queue.head.val

		user, ok := cache.Get(headId)
		tst.Equal(ok, true)
		tst.Equal(user, expUser)
		tst.Equal(cache.queue.size, uint(2))
		tst.Equal(cache.queue.tail.val, user)
	})

	t.Run("get tail user", func(t *testing.T) {
		headId := cache.queue.tail.key
		expUser := cache.queue.tail.val

		user, ok := cache.Get(headId)
		tst.Equal(ok, true)
		tst.Equal(user, expUser)
		tst.Equal(cache.queue.size, uint(2))
		tst.Equal(cache.queue.tail.val, user)
	})

}

func prepareSomeCache() *cache {
	cache := &cache{
		mx:       &sync.Mutex{},
		capacity: cap,
		users:    make(map[string]*node, cap),
		queue:    &linkedList{size: 0},
	}

	users := getSomeFunUsers(3)
	for _, user := range users {
		cache.Set(user)
	}

	return cache
}

func getSomeFunUsers(n int) []entities.User {
	users := make([]entities.User, 0)
	for i := 0; i < n; i++ {
		user := entities.User{
			ID:       fmt.Sprint(i),
			Nickname: fmt.Sprintf("crosh%d", i),
			Age:      uint(i),
		}
		users = append(users, user)
	}
	return users
}
