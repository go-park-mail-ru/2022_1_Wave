// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	albumProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"

	mock "github.com/stretchr/testify/mock"
)

// AlbumCoverRepo is an autogenerated mock type for the AlbumCoverRepo type
type AlbumCoverRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *AlbumCoverRepo) Create(_a0 *albumProto.AlbumCover) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*albumProto.AlbumCover) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *AlbumCoverRepo) Delete(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *AlbumCoverRepo) GetAll() ([]*albumProto.AlbumCover, error) {
	ret := _m.Called()

	var r0 []*albumProto.AlbumCover
	if rf, ok := ret.Get(0).(func() []*albumProto.AlbumCover); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*albumProto.AlbumCover)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastId provides a mock function with given fields:
func (_m *AlbumCoverRepo) GetLastId() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSize provides a mock function with given fields:
func (_m *AlbumCoverRepo) GetSize() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByID provides a mock function with given fields: _a0
func (_m *AlbumCoverRepo) SelectByID(_a0 int64) (*albumProto.AlbumCover, error) {
	ret := _m.Called(_a0)

	var r0 *albumProto.AlbumCover
	if rf, ok := ret.Get(0).(func(int64) *albumProto.AlbumCover); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*albumProto.AlbumCover)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *AlbumCoverRepo) Update(_a0 *albumProto.AlbumCover) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*albumProto.AlbumCover) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}