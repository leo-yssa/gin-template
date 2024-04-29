// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	dao "gin-api/app/domain/dao"

	mock "github.com/stretchr/testify/mock"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: role
func (_m *RoleRepository) Save(role *dao.Role) (dao.Role, error) {
	ret := _m.Called(role)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 dao.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(*dao.Role) (dao.Role, error)); ok {
		return rf(role)
	}
	if rf, ok := ret.Get(0).(func(*dao.Role) dao.Role); ok {
		r0 = rf(role)
	} else {
		r0 = ret.Get(0).(dao.Role)
	}

	if rf, ok := ret.Get(1).(func(*dao.Role) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRoleRepository creates a new instance of RoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoleRepository {
	mock := &RoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}