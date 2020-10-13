// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym/validator/nym/directory/models"
	mock "github.com/stretchr/testify/mock"
)

// GatewayHostSanitizer is an autogenerated mock type for the GatewayHostSanitizer type
type GatewayHostSanitizer struct {
	mock.Mock
}

// Sanitize provides a mock function with given fields: _a0
func (_m *GatewayHostSanitizer) Sanitize(_a0 models.GatewayHostInfo) models.GatewayHostInfo {
	ret := _m.Called(_a0)

	var r0 models.GatewayHostInfo
	if rf, ok := ret.Get(0).(func(models.GatewayHostInfo) models.GatewayHostInfo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.GatewayHostInfo)
	}

	return r0
}
