// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym/validator/nym/directory/models"
	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// CreateMixStatus provides a mock function with given fields: metric
func (_m *IService) CreateMixStatus(metric models.MixStatus) models.PersistedMixStatus {
	ret := _m.Called(metric)

	var r0 models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(models.MixStatus) models.PersistedMixStatus); ok {
		r0 = rf(metric)
	} else {
		r0 = ret.Get(0).(models.PersistedMixStatus)
	}

	return r0
}

// GetStatusReport provides a mock function with given fields: pubkey
func (_m *IService) GetStatusReport(pubkey string) models.MixStatusReport {
	ret := _m.Called(pubkey)

	var r0 models.MixStatusReport
	if rf, ok := ret.Get(0).(func(string) models.MixStatusReport); ok {
		r0 = rf(pubkey)
	} else {
		r0 = ret.Get(0).(models.MixStatusReport)
	}

	return r0
}

// List provides a mock function with given fields: pubkey
func (_m *IService) List(pubkey string) []models.PersistedMixStatus {
	ret := _m.Called(pubkey)

	var r0 []models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(string) []models.PersistedMixStatus); ok {
		r0 = rf(pubkey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixStatus)
		}
	}

	return r0
}

// SaveStatusReport provides a mock function with given fields: status
func (_m *IService) SaveStatusReport(status models.PersistedMixStatus) models.MixStatusReport {
	ret := _m.Called(status)

	var r0 models.MixStatusReport
	if rf, ok := ret.Get(0).(func(models.PersistedMixStatus) models.MixStatusReport); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(models.MixStatusReport)
	}

	return r0
}
