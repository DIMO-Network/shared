// Code generated by MockGen. DO NOT EDIT.
// Source: elastic_search_api_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	elastic "github.com/DIMO-Network/shared/elastic"
	gomock "github.com/golang/mock/gomock"
)

// MockAppSearchService is a mock of AppSearchService interface.
type MockAppSearchService struct {
	ctrl     *gomock.Controller
	recorder *MockAppSearchServiceMockRecorder
}

// MockAppSearchServiceMockRecorder is the mock recorder for MockAppSearchService.
type MockAppSearchServiceMockRecorder struct {
	mock *MockAppSearchService
}

// NewMockAppSearchService creates a new mock instance.
func NewMockAppSearchService(ctrl *gomock.Controller) *MockAppSearchService {
	mock := &MockAppSearchService{ctrl: ctrl}
	mock.recorder = &MockAppSearchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppSearchService) EXPECT() *MockAppSearchServiceMockRecorder {
	return m.recorder
}

// AddSourceEngineToMetaEngine mocks base method.
func (m *MockAppSearchService) AddSourceEngineToMetaEngine(sourceName, metaName string) (*elastic.EngineDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSourceEngineToMetaEngine", sourceName, metaName)
	ret0, _ := ret[0].(*elastic.EngineDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSourceEngineToMetaEngine indicates an expected call of AddSourceEngineToMetaEngine.
func (mr *MockAppSearchServiceMockRecorder) AddSourceEngineToMetaEngine(sourceName, metaName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSourceEngineToMetaEngine", reflect.TypeOf((*MockAppSearchService)(nil).AddSourceEngineToMetaEngine), sourceName, metaName)
}

// CreateDocuments mocks base method.
func (m *MockAppSearchService) CreateDocuments(docs []any, engineName string) ([]elastic.CreateDocsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDocuments", docs, engineName)
	ret0, _ := ret[0].([]elastic.CreateDocsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDocuments indicates an expected call of CreateDocuments.
func (mr *MockAppSearchServiceMockRecorder) CreateDocuments(docs, engineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDocuments", reflect.TypeOf((*MockAppSearchService)(nil).CreateDocuments), docs, engineName)
}

// CreateDocumentsBatched mocks base method.
func (m *MockAppSearchService) CreateDocumentsBatched(docs []any, engineName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDocumentsBatched", docs, engineName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDocumentsBatched indicates an expected call of CreateDocumentsBatched.
func (mr *MockAppSearchServiceMockRecorder) CreateDocumentsBatched(docs, engineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDocumentsBatched", reflect.TypeOf((*MockAppSearchService)(nil).CreateDocumentsBatched), docs, engineName)
}

// CreateEngine mocks base method.
func (m *MockAppSearchService) CreateEngine(name string, metaSource *string) (*elastic.EngineDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEngine", name, metaSource)
	ret0, _ := ret[0].(*elastic.EngineDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEngine indicates an expected call of CreateEngine.
func (mr *MockAppSearchServiceMockRecorder) CreateEngine(name, metaSource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEngine", reflect.TypeOf((*MockAppSearchService)(nil).CreateEngine), name, metaSource)
}

// DeleteEngine mocks base method.
func (m *MockAppSearchService) DeleteEngine(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEngine", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEngine indicates an expected call of DeleteEngine.
func (mr *MockAppSearchServiceMockRecorder) DeleteEngine(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEngine", reflect.TypeOf((*MockAppSearchService)(nil).DeleteEngine), name)
}

// GetEngines mocks base method.
func (m *MockAppSearchService) GetEngines() (*elastic.GetEnginesResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEngines")
	ret0, _ := ret[0].(*elastic.GetEnginesResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEngines indicates an expected call of GetEngines.
func (mr *MockAppSearchServiceMockRecorder) GetEngines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEngines", reflect.TypeOf((*MockAppSearchService)(nil).GetEngines))
}

// GetMetaEngineName mocks base method.
func (m *MockAppSearchService) GetMetaEngineName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaEngineName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMetaEngineName indicates an expected call of GetMetaEngineName.
func (mr *MockAppSearchServiceMockRecorder) GetMetaEngineName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaEngineName", reflect.TypeOf((*MockAppSearchService)(nil).GetMetaEngineName))
}

// RemoveSourceEngine mocks base method.
func (m *MockAppSearchService) RemoveSourceEngine(sourceName, metaName string) (*elastic.EngineDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSourceEngine", sourceName, metaName)
	ret0, _ := ret[0].(*elastic.EngineDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSourceEngine indicates an expected call of RemoveSourceEngine.
func (mr *MockAppSearchServiceMockRecorder) RemoveSourceEngine(sourceName, metaName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSourceEngine", reflect.TypeOf((*MockAppSearchService)(nil).RemoveSourceEngine), sourceName, metaName)
}