// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	artistProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	gatewayProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"

	mock "github.com/stretchr/testify/mock"
)

// ArtistAgent is an autogenerated mock type for the ArtistAgent type
type ArtistAgent struct {
	mock.Mock
}

// AddToFavorites provides a mock function with given fields: data
func (_m *ArtistAgent) AddToFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	ret := _m.Called(data)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(*gatewayProto.UserIdArtistIdArg) *emptypb.Empty); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gatewayProto.UserIdArtistIdArg) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0
func (_m *ArtistAgent) Create(_a0 *artistProto.Artist) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*artistProto.Artist) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *ArtistAgent) Delete(_a0 *gatewayProto.IdArg) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gatewayProto.IdArg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ArtistAgent) GetAll() (*artistProto.ArtistsResponse, error) {
	ret := _m.Called()

	var r0 *artistProto.ArtistsResponse
	if rf, ok := ret.Get(0).(func() *artistProto.ArtistsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artistProto.ArtistsResponse)
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

// GetById provides a mock function with given fields: _a0
func (_m *ArtistAgent) GetById(_a0 *gatewayProto.IdArg) (*artistProto.ArtistDataTransfer, error) {
	ret := _m.Called(_a0)

	var r0 *artistProto.ArtistDataTransfer
	if rf, ok := ret.Get(0).(func(*gatewayProto.IdArg) *artistProto.ArtistDataTransfer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artistProto.ArtistDataTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gatewayProto.IdArg) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFavorites provides a mock function with given fields: _a0
func (_m *ArtistAgent) GetFavorites(_a0 *gatewayProto.IdArg) (*artistProto.ArtistsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *artistProto.ArtistsResponse
	if rf, ok := ret.Get(0).(func(*gatewayProto.IdArg) *artistProto.ArtistsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artistProto.ArtistsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gatewayProto.IdArg) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastId provides a mock function with given fields:
func (_m *ArtistAgent) GetLastId() (*gatewayProto.IntResponse, error) {
	ret := _m.Called()

	var r0 *gatewayProto.IntResponse
	if rf, ok := ret.Get(0).(func() *gatewayProto.IntResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gatewayProto.IntResponse)
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

// GetPopular provides a mock function with given fields:
func (_m *ArtistAgent) GetPopular() (*artistProto.ArtistsResponse, error) {
	ret := _m.Called()

	var r0 *artistProto.ArtistsResponse
	if rf, ok := ret.Get(0).(func() *artistProto.ArtistsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artistProto.ArtistsResponse)
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

// GetSize provides a mock function with given fields:
func (_m *ArtistAgent) GetSize() (*gatewayProto.IntResponse, error) {
	ret := _m.Called()

	var r0 *gatewayProto.IntResponse
	if rf, ok := ret.Get(0).(func() *gatewayProto.IntResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gatewayProto.IntResponse)
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

// RemoveFromFavorites provides a mock function with given fields: data
func (_m *ArtistAgent) RemoveFromFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	ret := _m.Called(data)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(*gatewayProto.UserIdArtistIdArg) *emptypb.Empty); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gatewayProto.UserIdArtistIdArg) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchByName provides a mock function with given fields: arg
func (_m *ArtistAgent) SearchByName(arg *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error) {
	ret := _m.Called(arg)

	var r0 *artistProto.ArtistsResponse
	if rf, ok := ret.Get(0).(func(*gatewayProto.StringArg) *artistProto.ArtistsResponse); ok {
		r0 = rf(arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artistProto.ArtistsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gatewayProto.StringArg) error); ok {
		r1 = rf(arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *ArtistAgent) Update(_a0 *artistProto.Artist) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*artistProto.Artist) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}