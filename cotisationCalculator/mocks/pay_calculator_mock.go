// Code generated by MockGen. DO NOT EDIT.
// Source: paycalculator/pay_calculator.go

// Package mock_paycalculator is a generated GoMock package.
package mock_paycalculator

import (
	data "cotisationCalculator/data"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPayDataProvider is a mock of PayDataProvider interface.
type MockPayDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPayDataProviderMockRecorder
}

// MockPayDataProviderMockRecorder is the mock recorder for MockPayDataProvider.
type MockPayDataProviderMockRecorder struct {
	mock *MockPayDataProvider
}

// NewMockPayDataProvider creates a new mock instance.
func NewMockPayDataProvider(ctrl *gomock.Controller) *MockPayDataProvider {
	mock := &MockPayDataProvider{ctrl: ctrl}
	mock.recorder = &MockPayDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPayDataProvider) EXPECT() *MockPayDataProviderMockRecorder {
	return m.recorder
}

// GetCotisation mocks base method.
func (m *MockPayDataProvider) GetCotisation(cotisation string, infoEntreprise data.InfoEntreprise, salaire float32) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCotisation", cotisation, infoEntreprise, salaire)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCotisation indicates an expected call of GetCotisation.
func (mr *MockPayDataProviderMockRecorder) GetCotisation(cotisation, infoEntreprise, salaire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCotisation", reflect.TypeOf((*MockPayDataProvider)(nil).GetCotisation), cotisation, infoEntreprise, salaire)
}

// MockCotisation is a mock of Cotisation interface.
type MockCotisation struct {
	ctrl     *gomock.Controller
	recorder *MockCotisationMockRecorder
}

// MockCotisationMockRecorder is the mock recorder for MockCotisation.
type MockCotisationMockRecorder struct {
	mock *MockCotisation
}

// NewMockCotisation creates a new mock instance.
func NewMockCotisation(ctrl *gomock.Controller) *MockCotisation {
	mock := &MockCotisation{ctrl: ctrl}
	mock.recorder = &MockCotisationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCotisation) EXPECT() *MockCotisationMockRecorder {
	return m.recorder
}

// ToString mocks base method.
func (m *MockCotisation) ToString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToString")
	ret0, _ := ret[0].(string)
	return ret0
}

// ToString indicates an expected call of ToString.
func (mr *MockCotisationMockRecorder) ToString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToString", reflect.TypeOf((*MockCotisation)(nil).ToString))
}

// ToUrssaf mocks base method.
func (m *MockCotisation) ToUrssaf() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUrssaf")
	ret0, _ := ret[0].(string)
	return ret0
}

// ToUrssaf indicates an expected call of ToUrssaf.
func (mr *MockCotisationMockRecorder) ToUrssaf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUrssaf", reflect.TypeOf((*MockCotisation)(nil).ToUrssaf))
}

// MockCotisationProvider is a mock of CotisationProvider interface.
type MockCotisationProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCotisationProviderMockRecorder
}

// MockCotisationProviderMockRecorder is the mock recorder for MockCotisationProvider.
type MockCotisationProviderMockRecorder struct {
	mock *MockCotisationProvider
}

// NewMockCotisationProvider creates a new mock instance.
func NewMockCotisationProvider(ctrl *gomock.Controller) *MockCotisationProvider {
	mock := &MockCotisationProvider{ctrl: ctrl}
	mock.recorder = &MockCotisationProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCotisationProvider) EXPECT() *MockCotisationProviderMockRecorder {
	return m.recorder
}

// GetCotisation mocks base method.
func (m *MockCotisationProvider) GetCotisation(cotisation string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCotisation", cotisation)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCotisation indicates an expected call of GetCotisation.
func (mr *MockCotisationProviderMockRecorder) GetCotisation(cotisation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCotisation", reflect.TypeOf((*MockCotisationProvider)(nil).GetCotisation), cotisation)
}

// MockPayCotisationsInterface is a mock of PayCotisationsInterface interface.
type MockPayCotisationsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPayCotisationsInterfaceMockRecorder
}

// MockPayCotisationsInterfaceMockRecorder is the mock recorder for MockPayCotisationsInterface.
type MockPayCotisationsInterfaceMockRecorder struct {
	mock *MockPayCotisationsInterface
}

// NewMockPayCotisationsInterface creates a new mock instance.
func NewMockPayCotisationsInterface(ctrl *gomock.Controller) *MockPayCotisationsInterface {
	mock := &MockPayCotisationsInterface{ctrl: ctrl}
	mock.recorder = &MockPayCotisationsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPayCotisationsInterface) EXPECT() *MockPayCotisationsInterfaceMockRecorder {
	return m.recorder
}

// CotisationPatronaleForAPI mocks base method.
func (m *MockPayCotisationsInterface) CotisationPatronaleForAPI(salaire int) map[string]float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CotisationPatronaleForAPI", salaire)
	ret0, _ := ret[0].(map[string]float64)
	return ret0
}

// CotisationPatronaleForAPI indicates an expected call of CotisationPatronaleForAPI.
func (mr *MockPayCotisationsInterfaceMockRecorder) CotisationPatronaleForAPI(salaire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CotisationPatronaleForAPI", reflect.TypeOf((*MockPayCotisationsInterface)(nil).CotisationPatronaleForAPI), salaire)
}
