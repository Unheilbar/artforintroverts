package user

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

func TestService_TestCacheInvalidation(t *testing.T) {

	t.Run("valid cache no storage calls", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		sp.EXPECT().Cache().Return(cache)
		cache.EXPECT().IsValid().Return(true)
		cache.EXPECT().GetAll().Return([]entities.User{})

		svc := NewService(sp)
		_, err := svc.GetAll()
		tst.NoError(err)
	})

	t.Run("invalid cache storage calls no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)
		sp.EXPECT().Cache().Return(cache)
		cache.EXPECT().IsValid().Return(false)
		sp.EXPECT().Storage().Return(st)
		st.EXPECT().GetAll()

		svc := NewService(sp)
		_, err := svc.GetAll()
		tst.NoError(err)
	})

	t.Run("invalidate cache after update event with no errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)

		sp.EXPECT().Storage().Return(st)
		sp.EXPECT().Cache().Return(cache)
		st.EXPECT().Delete(gomock.Any())
		cache.EXPECT().SetInvalid()

		svc := NewService(sp)
		err := svc.Delete("123")
		tst.NoError(err)
	})

	t.Run("cancel invalidate cache after update event with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)

		sp.EXPECT().Storage().Return(st)
		sp.EXPECT().Cache().Return(cache).Times(2)
		deleteErr := errors.New("delete event failed")
		st.EXPECT().Delete(gomock.Any()).Return(deleteErr)
		cache.EXPECT().SetInvalid()
		cache.EXPECT().SetValid()

		svc := NewService(sp)
		err := svc.Delete("123")
		tst.Equal(deleteErr, err)
	})
}
