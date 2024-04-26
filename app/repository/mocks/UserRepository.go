// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	dao "gin-api/app/domain/dao"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: user
func (_m *UserRepository) Save(user *dao.User) (dao.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 dao.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*dao.User) (dao.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*dao.User) dao.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(dao.User)
	}

	if rf, ok := ret.Get(1).(func(*dao.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
