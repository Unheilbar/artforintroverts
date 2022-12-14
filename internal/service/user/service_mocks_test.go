// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user/service.go

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStorage) Delete(ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorage)(nil).Delete), ID)
}

// GetAll mocks base method.
func (m *MockStorage) GetAll() ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStorageMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStorage)(nil).GetAll))
}

// Update mocks base method.
func (m *MockStorage) Update(ID string, user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStorageMockRecorder) Update(ID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStorage)(nil).Update), ID, user)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockCache) GetAll() []entities.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.User)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCacheMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCache)(nil).GetAll))
}

// IsValid mocks base method.
func (m *MockCache) IsValid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockCacheMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockCache)(nil).IsValid))
}

// SetInvalid mocks base method.
func (m *MockCache) SetInvalid() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInvalid")
}

// SetInvalid indicates an expected call of SetInvalid.
func (mr *MockCacheMockRecorder) SetInvalid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInvalid", reflect.TypeOf((*MockCache)(nil).SetInvalid))
}

// SetValid mocks base method.
func (m *MockCache) SetValid() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValid")
}

// SetValid indicates an expected call of SetValid.
func (mr *MockCacheMockRecorder) SetValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValid", reflect.TypeOf((*MockCache)(nil).SetValid))
}

// MockStorageProvider is a mock of StorageProvider interface.
type MockStorageProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStorageProviderMockRecorder
}

// MockStorageProviderMockRecorder is the mock recorder for MockStorageProvider.
type MockStorageProviderMockRecorder struct {
	mock *MockStorageProvider
}

// NewMockStorageProvider creates a new mock instance.
func NewMockStorageProvider(ctrl *gomock.Controller) *MockStorageProvider {
	mock := &MockStorageProvider{ctrl: ctrl}
	mock.recorder = &MockStorageProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageProvider) EXPECT() *MockStorageProviderMockRecorder {
	return m.recorder
}

// Cache mocks base method.
func (m *MockStorageProvider) Cache() Cache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cache")
	ret0, _ := ret[0].(Cache)
	return ret0
}

// Cache indicates an expected call of Cache.
func (mr *MockStorageProviderMockRecorder) Cache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cache", reflect.TypeOf((*MockStorageProvider)(nil).Cache))
}

// Storage mocks base method.
func (m *MockStorageProvider) Storage() Storage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(Storage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockStorageProviderMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockStorageProvider)(nil).Storage))
}
