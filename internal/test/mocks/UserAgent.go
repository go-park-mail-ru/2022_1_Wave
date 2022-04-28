// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
)

// UserAgent is an autogenerated mock type for the UserAgent type
type UserAgent struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserAgent) Create(user *user_microservice_domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user_microservice_domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *UserAgent) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserAgent) GetByEmail(email string) (*user_microservice_domain.User, error) {
	ret := _m.Called(email)

	var r0 *user_microservice_domain.User
	if rf, ok := ret.Get(0).(func(string) *user_microservice_domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user_microservice_domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *UserAgent) GetById(id uint) (*user_microservice_domain.User, error) {
	ret := _m.Called(id)

	var r0 *user_microservice_domain.User
	if rf, ok := ret.Get(0).(func(uint) *user_microservice_domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user_microservice_domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *UserAgent) GetByUsername(username string) (*user_microservice_domain.User, error) {
	ret := _m.Called(username)

	var r0 *user_microservice_domain.User
	if rf, ok := ret.Get(0).(func(string) *user_microservice_domain.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user_microservice_domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, user
func (_m *UserAgent) Update(id uint, user *user_microservice_domain.User) error {
	ret := _m.Called(id, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *user_microservice_domain.User) error); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
