// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app "github.com/reckycless/event-checker"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user app.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), login, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockEvent is a mock of Event interface.
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
}

// MockEventMockRecorder is the mock recorder for MockEvent.
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance.
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEvent) Create(userID int, event app.Event) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, event)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEventMockRecorder) Create(userID, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEvent)(nil).Create), userID, event)
}

// Delete mocks base method.
func (m *MockEvent) Delete(userID, eventID int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, eventID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEventMockRecorder) Delete(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEvent)(nil).Delete), userID, eventID)
}

// GetAll mocks base method.
func (m *MockEvent) GetAll() ([]app.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]app.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEvent)(nil).GetAll))
}

// GetAllForUser mocks base method.
func (m *MockEvent) GetAllForUser(userID int) ([]app.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllForUser", userID)
	ret0, _ := ret[0].([]app.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllForUser indicates an expected call of GetAllForUser.
func (mr *MockEventMockRecorder) GetAllForUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllForUser", reflect.TypeOf((*MockEvent)(nil).GetAllForUser), userID)
}

// GetByID mocks base method.
func (m *MockEvent) GetByID(eventID int) (app.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", eventID)
	ret0, _ := ret[0].(app.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEventMockRecorder) GetByID(eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEvent)(nil).GetByID), eventID)
}

// Update mocks base method.
func (m *MockEvent) Update(userID, eventID int, input app.UpdateEventInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, eventID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEventMockRecorder) Update(userID, eventID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEvent)(nil).Update), userID, eventID, input)
}

// MockVisitor is a mock of Visitor interface.
type MockVisitor struct {
	ctrl     *gomock.Controller
	recorder *MockVisitorMockRecorder
}

// MockVisitorMockRecorder is the mock recorder for MockVisitor.
type MockVisitorMockRecorder struct {
	mock *MockVisitor
}

// NewMockVisitor creates a new mock instance.
func NewMockVisitor(ctrl *gomock.Controller) *MockVisitor {
	mock := &MockVisitor{ctrl: ctrl}
	mock.recorder = &MockVisitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisitor) EXPECT() *MockVisitorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVisitor) Create(eventID int, visitor app.Visitor) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", eventID, visitor)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVisitorMockRecorder) Create(eventID, visitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVisitor)(nil).Create), eventID, visitor)
}

// CreateAsUser mocks base method.
func (m *MockVisitor) CreateAsUser(userID, eventID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAsUser", userID, eventID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAsUser indicates an expected call of CreateAsUser.
func (mr *MockVisitorMockRecorder) CreateAsUser(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAsUser", reflect.TypeOf((*MockVisitor)(nil).CreateAsUser), userID, eventID)
}

// Delete mocks base method.
func (m *MockVisitor) Delete(userID, visitorID int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, visitorID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockVisitorMockRecorder) Delete(userID, visitorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVisitor)(nil).Delete), userID, visitorID)
}

// GetAllEventVisitors mocks base method.
func (m *MockVisitor) GetAllEventVisitors(userID, eventID int) ([]app.Visitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEventVisitors", userID, eventID)
	ret0, _ := ret[0].([]app.Visitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEventVisitors indicates an expected call of GetAllEventVisitors.
func (mr *MockVisitorMockRecorder) GetAllEventVisitors(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEventVisitors", reflect.TypeOf((*MockVisitor)(nil).GetAllEventVisitors), userID, eventID)
}

// GetByID mocks base method.
func (m *MockVisitor) GetByID(userID, visitorID int) (app.Visitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", userID, visitorID)
	ret0, _ := ret[0].(app.Visitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockVisitorMockRecorder) GetByID(userID, visitorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVisitor)(nil).GetByID), userID, visitorID)
}

// Update mocks base method.
func (m *MockVisitor) Update(userID, visitorID int, input app.UpdateVisitorInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, visitorID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVisitorMockRecorder) Update(userID, visitorID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVisitor)(nil).Update), userID, visitorID, input)
}

// MockOrganisator is a mock of Organisator interface.
type MockOrganisator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganisatorMockRecorder
}

// MockOrganisatorMockRecorder is the mock recorder for MockOrganisator.
type MockOrganisatorMockRecorder struct {
	mock *MockOrganisator
}

// NewMockOrganisator creates a new mock instance.
func NewMockOrganisator(ctrl *gomock.Controller) *MockOrganisator {
	mock := &MockOrganisator{ctrl: ctrl}
	mock.recorder = &MockOrganisatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganisator) EXPECT() *MockOrganisatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganisator) Create(visitor app.Organisator) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", visitor)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganisatorMockRecorder) Create(visitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganisator)(nil).Create), visitor)
}

// Delete mocks base method.
func (m *MockOrganisator) Delete(organisatorID int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", organisatorID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganisatorMockRecorder) Delete(organisatorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganisator)(nil).Delete), organisatorID)
}

// GetAll mocks base method.
func (m *MockOrganisator) GetAll() ([]app.Organisator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]app.Organisator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOrganisatorMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOrganisator)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockOrganisator) GetByID(organisatorID int) (app.Organisator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", organisatorID)
	ret0, _ := ret[0].(app.Organisator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrganisatorMockRecorder) GetByID(organisatorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrganisator)(nil).GetByID), organisatorID)
}

// Update mocks base method.
func (m *MockOrganisator) Update(organisatorID int, input app.UpdateOrganisatorInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", organisatorID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganisatorMockRecorder) Update(organisatorID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganisator)(nil).Update), organisatorID, input)
}

// MockEventType is a mock of EventType interface.
type MockEventType struct {
	ctrl     *gomock.Controller
	recorder *MockEventTypeMockRecorder
}

// MockEventTypeMockRecorder is the mock recorder for MockEventType.
type MockEventTypeMockRecorder struct {
	mock *MockEventType
}

// NewMockEventType creates a new mock instance.
func NewMockEventType(ctrl *gomock.Controller) *MockEventType {
	mock := &MockEventType{ctrl: ctrl}
	mock.recorder = &MockEventTypeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventType) EXPECT() *MockEventTypeMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEventType) Create(eventType app.EventType) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", eventType)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEventTypeMockRecorder) Create(eventType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEventType)(nil).Create), eventType)
}

// Delete mocks base method.
func (m *MockEventType) Delete(eventTypeID int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", eventTypeID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEventTypeMockRecorder) Delete(eventTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventType)(nil).Delete), eventTypeID)
}

// GetAll mocks base method.
func (m *MockEventType) GetAll() ([]app.EventType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]app.EventType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventTypeMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEventType)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockEventType) GetByID(eventTypeID int) (app.EventType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", eventTypeID)
	ret0, _ := ret[0].(app.EventType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEventTypeMockRecorder) GetByID(eventTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEventType)(nil).GetByID), eventTypeID)
}

// Update mocks base method.
func (m *MockEventType) Update(eventTypeID int, input app.UpdateEventTypeInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", eventTypeID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEventTypeMockRecorder) Update(eventTypeID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventType)(nil).Update), eventTypeID, input)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// UpdateUserOrganisator mocks base method.
func (m *MockUser) UpdateUserOrganisator(userID int, input app.UserOrganisatorInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserOrganisator", userID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserOrganisator indicates an expected call of UpdateUserOrganisator.
func (mr *MockUserMockRecorder) UpdateUserOrganisator(userID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserOrganisator", reflect.TypeOf((*MockUser)(nil).UpdateUserOrganisator), userID, input)
}

// UpdateUserRole mocks base method.
func (m *MockUser) UpdateUserRole(userID int, input app.UserRoleInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRole", userID, input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserRole indicates an expected call of UpdateUserRole.
func (mr *MockUserMockRecorder) UpdateUserRole(userID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRole", reflect.TypeOf((*MockUser)(nil).UpdateUserRole), userID, input)
}