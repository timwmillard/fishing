// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/timwmillard/fishing (interfaces: CompetitorsRepo)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	fishing "github.com/timwmillard/fishing"
)

// CompetitorsRepo is a mock of CompetitorsRepo interface.
type CompetitorsRepo struct {
	ctrl     *gomock.Controller
	recorder *CompetitorsRepoMockRecorder
}

// CompetitorsRepoMockRecorder is the mock recorder for CompetitorsRepo.
type CompetitorsRepoMockRecorder struct {
	mock *CompetitorsRepo
}

// NewCompetitorsRepo creates a new mock instance.
func NewCompetitorsRepo(ctrl *gomock.Controller) *CompetitorsRepo {
	mock := &CompetitorsRepo{ctrl: ctrl}
	mock.recorder = &CompetitorsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CompetitorsRepo) EXPECT() *CompetitorsRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *CompetitorsRepo) Create(arg0 context.Context, arg1 fishing.Competitor) (fishing.Competitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(fishing.Competitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *CompetitorsRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CompetitorsRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *CompetitorsRepo) Delete(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *CompetitorsRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*CompetitorsRepo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *CompetitorsRepo) Get(arg0 context.Context, arg1 uuid.UUID) (fishing.Competitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(fishing.Competitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *CompetitorsRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*CompetitorsRepo)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *CompetitorsRepo) List(arg0 context.Context) ([]fishing.Competitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]fishing.Competitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *CompetitorsRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*CompetitorsRepo)(nil).List), arg0)
}

// Update mocks base method.
func (m *CompetitorsRepo) Update(arg0 context.Context, arg1 fishing.Competitor) (fishing.Competitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(fishing.Competitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *CompetitorsRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*CompetitorsRepo)(nil).Update), arg0, arg1)
}
