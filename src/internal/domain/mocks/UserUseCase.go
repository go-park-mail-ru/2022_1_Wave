// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// CheckEmailAndPassword provides a mock function with given fields: email, password
func (_m *UserUseCase) CheckEmailAndPassword(email string, password string) bool {
	ret := _m.Called(email, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CheckUsernameAndPassword provides a mock function with given fields: username, password
func (_m *UserUseCase) CheckUsernameAndPassword(username string, password string) bool {
	ret := _m.Called(username, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteByEmail provides a mock function with given fields: email
func (_m *UserUseCase) DeleteByEmail(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: userId
func (_m *UserUseCase) DeleteById(userId uint) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBySessionId provides a mock function with given fields: sessionId
func (_m *UserUseCase) DeleteBySessionId(sessionId string) error {
	ret := _m.Called(sessionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByUsername provides a mock function with given fields: username
func (_m *UserUseCase) DeleteByUsername(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserUseCase) GetByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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

// GetById provides a mock function with given fields: userId
func (_m *UserUseCase) GetById(userId uint) (*domain.User, error) {
	ret := _m.Called(userId)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(uint) *domain.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySessionId provides a mock function with given fields: sessionId
func (_m *UserUseCase) GetBySessionId(sessionId string) (*domain.User, error) {
	ret := _m.Called(sessionId)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(sessionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *UserUseCase) GetByUsername(username string) (*domain.User, error) {
	ret := _m.Called(username)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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
func (_m *UserUseCase) Update(id uint, user *domain.User) error {
	ret := _m.Called(id, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *domain.User) error); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
