package user

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

func TestService_TestCacheInvalidation(t *testing.T) {
	t.Run("delete call triggers delete cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)

		sp.EXPECT().Storage().Return(st)
		sp.EXPECT().Cache().Return(cache)
		st.EXPECT().Delete(gomock.Any())
		cache.EXPECT().Delete(gomock.Any())

		svc := NewService(sp)
		err := svc.Delete("123")
		tst.NoError(err)
	})

	t.Run("update call triggers set cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)

		sp.EXPECT().Storage().Return(st)
		sp.EXPECT().Cache().Return(cache)
		st.EXPECT().UpdateAndReturn(gomock.Any(), gomock.Any())
		cache.EXPECT().Set(gomock.Any())

		svc := NewService(sp)
		err := svc.Update("123", entities.User{})
		tst.NoError(err)
	})

	t.Run("get by id triggers cache set in case of cache miss", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tst := assert.New(t)

		sp := NewMockStorageProvider(ctrl)
		cache := NewMockCache(ctrl)
		st := NewMockStorage(ctrl)

		cache.EXPECT().Get(gomock.Any())
		sp.EXPECT().Storage().Return(st)
		sp.EXPECT().Cache().Return(cache).Times(2)
		st.EXPECT().GetById(gomock.Any())
		cache.EXPECT().Set(gomock.Any())

		svc := NewService(sp)
		_, err := svc.GetById("123")
		tst.NoError(err)
	})
}
