// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/rmfachran/miniproject2/entity"
	mock "github.com/stretchr/testify/mock"
)

// ActorInterfaceRepo is an autogenerated mock type for the ActorInterfaceRepo type
type ActorInterfaceRepo struct {
	mock.Mock
}

// ApprovedAdmin provides a mock function with given fields: id
func (_m *ActorInterfaceRepo) ApprovedAdmin(id uint) ([]*entity.Actor, error) {
	ret := _m.Called(id)

	var r0 []*entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]*entity.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) []*entity.Actor); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdmin provides a mock function with given fields: actor
func (_m *ActorInterfaceRepo) CreateAdmin(actor *entity.Actor) (*entity.Actor, error) {
	ret := _m.Called(actor)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) (*entity.Actor, error)); ok {
		return rf(actor)
	}
	if rf, ok := ret.Get(0).(func(*entity.Actor) *entity.Actor); ok {
		r0 = rf(actor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Actor) error); ok {
		r1 = rf(actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAdmin provides a mock function with given fields: id, actor
func (_m *ActorInterfaceRepo) DeleteAdmin(id uint, actor *entity.Actor) error {
	ret := _m.Called(id, actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *entity.Actor) error); ok {
		r0 = rf(id, actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdmin provides a mock function with given fields: id
func (_m *ActorInterfaceRepo) GetAdmin(id uint) (*entity.Actor, error) {
	ret := _m.Called(id)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entity.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entity.Actor); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomers provides a mock function with given fields: first_name, last_name, email, page, pageSize
func (_m *ActorInterfaceRepo) GetCustomers(first_name string, last_name string, email string, page int, pageSize int) ([]*entity.Customer, error) {
	ret := _m.Called(first_name, last_name, email, page, pageSize)

	var r0 []*entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) ([]*entity.Customer, error)); ok {
		return rf(first_name, last_name, email, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) []*entity.Customer); ok {
		r0 = rf(first_name, last_name, email, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int, int) error); ok {
		r1 = rf(first_name, last_name, email, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginAdmin provides a mock function with given fields: username
func (_m *ActorInterfaceRepo) LoginAdmin(username string) (*entity.Actor, error) {
	ret := _m.Called(username)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Actor); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginSuperAdmin provides a mock function with given fields: username
func (_m *ActorInterfaceRepo) LoginSuperAdmin(username string) (*entity.Actor, error) {
	ret := _m.Called(username)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Actor); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCustomersFromAPI provides a mock function with given fields: url
func (_m *ActorInterfaceRepo) SaveCustomersFromAPI(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdmin provides a mock function with given fields: id, actor
func (_m *ActorInterfaceRepo) UpdateAdmin(id uint, actor *entity.Actor) (*entity.Actor, error) {
	ret := _m.Called(id, actor)

	var r0 *entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *entity.Actor) (*entity.Actor, error)); ok {
		return rf(id, actor)
	}
	if rf, ok := ret.Get(0).(func(uint, *entity.Actor) *entity.Actor); ok {
		r0 = rf(id, actor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, *entity.Actor) error); ok {
		r1 = rf(id, actor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewActorInterfaceRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewActorInterfaceRepo creates a new instance of ActorInterfaceRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActorInterfaceRepo(t mockConstructorTestingTNewActorInterfaceRepo) *ActorInterfaceRepo {
	mock := &ActorInterfaceRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
