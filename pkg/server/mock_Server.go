// Code generated by mockery v2.46.3. DO NOT EDIT.

package server

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockServer is an autogenerated mock type for the Server type
type MockServer struct {
	mock.Mock
}

type MockServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServer) EXPECT() *MockServer_Expecter {
	return &MockServer_Expecter{mock: &_m.Mock}
}

// DELETE provides a mock function with given fields: path, h
func (_m *MockServer) DELETE(path string, h echo.HandlerFunc) {
	_m.Called(path, h)
}

// MockServer_DELETE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DELETE'
type MockServer_DELETE_Call struct {
	*mock.Call
}

// DELETE is a helper method to define mock.On call
//   - path string
//   - h echo.HandlerFunc
func (_e *MockServer_Expecter) DELETE(path interface{}, h interface{}) *MockServer_DELETE_Call {
	return &MockServer_DELETE_Call{Call: _e.mock.On("DELETE", path, h)}
}

func (_c *MockServer_DELETE_Call) Run(run func(path string, h echo.HandlerFunc)) *MockServer_DELETE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(echo.HandlerFunc))
	})
	return _c
}

func (_c *MockServer_DELETE_Call) Return() *MockServer_DELETE_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServer_DELETE_Call) RunAndReturn(run func(string, echo.HandlerFunc)) *MockServer_DELETE_Call {
	_c.Call.Return(run)
	return _c
}

// GET provides a mock function with given fields: path, h
func (_m *MockServer) GET(path string, h echo.HandlerFunc) {
	_m.Called(path, h)
}

// MockServer_GET_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GET'
type MockServer_GET_Call struct {
	*mock.Call
}

// GET is a helper method to define mock.On call
//   - path string
//   - h echo.HandlerFunc
func (_e *MockServer_Expecter) GET(path interface{}, h interface{}) *MockServer_GET_Call {
	return &MockServer_GET_Call{Call: _e.mock.On("GET", path, h)}
}

func (_c *MockServer_GET_Call) Run(run func(path string, h echo.HandlerFunc)) *MockServer_GET_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(echo.HandlerFunc))
	})
	return _c
}

func (_c *MockServer_GET_Call) Return() *MockServer_GET_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServer_GET_Call) RunAndReturn(run func(string, echo.HandlerFunc)) *MockServer_GET_Call {
	_c.Call.Return(run)
	return _c
}

// PATCH provides a mock function with given fields: path, h
func (_m *MockServer) PATCH(path string, h echo.HandlerFunc) {
	_m.Called(path, h)
}

// MockServer_PATCH_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PATCH'
type MockServer_PATCH_Call struct {
	*mock.Call
}

// PATCH is a helper method to define mock.On call
//   - path string
//   - h echo.HandlerFunc
func (_e *MockServer_Expecter) PATCH(path interface{}, h interface{}) *MockServer_PATCH_Call {
	return &MockServer_PATCH_Call{Call: _e.mock.On("PATCH", path, h)}
}

func (_c *MockServer_PATCH_Call) Run(run func(path string, h echo.HandlerFunc)) *MockServer_PATCH_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(echo.HandlerFunc))
	})
	return _c
}

func (_c *MockServer_PATCH_Call) Return() *MockServer_PATCH_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServer_PATCH_Call) RunAndReturn(run func(string, echo.HandlerFunc)) *MockServer_PATCH_Call {
	_c.Call.Return(run)
	return _c
}

// POST provides a mock function with given fields: path, h
func (_m *MockServer) POST(path string, h echo.HandlerFunc) {
	_m.Called(path, h)
}

// MockServer_POST_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'POST'
type MockServer_POST_Call struct {
	*mock.Call
}

// POST is a helper method to define mock.On call
//   - path string
//   - h echo.HandlerFunc
func (_e *MockServer_Expecter) POST(path interface{}, h interface{}) *MockServer_POST_Call {
	return &MockServer_POST_Call{Call: _e.mock.On("POST", path, h)}
}

func (_c *MockServer_POST_Call) Run(run func(path string, h echo.HandlerFunc)) *MockServer_POST_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(echo.HandlerFunc))
	})
	return _c
}

func (_c *MockServer_POST_Call) Return() *MockServer_POST_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServer_POST_Call) RunAndReturn(run func(string, echo.HandlerFunc)) *MockServer_POST_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: port
func (_m *MockServer) Start(port int) error {
	ret := _m.Called(port)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockServer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - port int
func (_e *MockServer_Expecter) Start(port interface{}) *MockServer_Start_Call {
	return &MockServer_Start_Call{Call: _e.mock.On("Start", port)}
}

func (_c *MockServer_Start_Call) Run(run func(port int)) *MockServer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockServer_Start_Call) Return(_a0 error) *MockServer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServer_Start_Call) RunAndReturn(run func(int) error) *MockServer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServer creates a new instance of MockServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServer {
	mock := &MockServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
