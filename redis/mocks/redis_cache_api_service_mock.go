// Code generated by MockGen. DO NOT EDIT.
// Source: redis_cache_api_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisCacheService is a mock of RedisCacheService interface.
type MockRedisCacheService struct {
	ctrl     *gomock.Controller
	recorder *MockRedisCacheServiceMockRecorder
}

// MockRedisCacheServiceMockRecorder is the mock recorder for MockRedisCacheService.
type MockRedisCacheServiceMockRecorder struct {
	mock *MockRedisCacheService
}

// NewMockRedisCacheService creates a new mock instance.
func NewMockRedisCacheService(ctrl *gomock.Controller) *MockRedisCacheService {
	mock := &MockRedisCacheService{ctrl: ctrl}
	mock.recorder = &MockRedisCacheServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisCacheService) EXPECT() *MockRedisCacheServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisCacheService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisCacheServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisCacheService)(nil).Close))
}

// Del mocks base method.
func (m *MockRedisCacheService) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisCacheServiceMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisCacheService)(nil).Del), varargs...)
}

// FlushAll mocks base method.
func (m *MockRedisCacheService) FlushAll(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockRedisCacheServiceMockRecorder) FlushAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockRedisCacheService)(nil).FlushAll), ctx)
}

// Get mocks base method.
func (m *MockRedisCacheService) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisCacheServiceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisCacheService)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockRedisCacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisCacheServiceMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisCacheService)(nil).Set), ctx, key, value, expiration)
}
