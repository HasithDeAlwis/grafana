// Code generated by mockery v2.53.4. DO NOT EDIT.

package pullrequest

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCommenter is an autogenerated mock type for the Commenter type
type MockCommenter struct {
	mock.Mock
}

type MockCommenter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommenter) EXPECT() *MockCommenter_Expecter {
	return &MockCommenter_Expecter{mock: &_m.Mock}
}

// Comment provides a mock function with given fields: ctx, repo, pr, changeInfo3
func (_m *MockCommenter) Comment(ctx context.Context, repo PullRequestRepo, pr int, changeInfo3 changeInfo) error {
	ret := _m.Called(ctx, repo, pr, changeInfo3)

	if len(ret) == 0 {
		panic("no return value specified for Comment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, PullRequestRepo, int, changeInfo) error); ok {
		r0 = rf(ctx, repo, pr, changeInfo3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommenter_Comment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Comment'
type MockCommenter_Comment_Call struct {
	*mock.Call
}

// Comment is a helper method to define mock.On call
//   - ctx context.Context
//   - repo PullRequestRepo
//   - pr int
//   - changeInfo3 changeInfo
func (_e *MockCommenter_Expecter) Comment(ctx interface{}, repo interface{}, pr interface{}, changeInfo3 interface{}) *MockCommenter_Comment_Call {
	return &MockCommenter_Comment_Call{Call: _e.mock.On("Comment", ctx, repo, pr, changeInfo3)}
}

func (_c *MockCommenter_Comment_Call) Run(run func(ctx context.Context, repo PullRequestRepo, pr int, changeInfo3 changeInfo)) *MockCommenter_Comment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(PullRequestRepo), args[2].(int), args[3].(changeInfo))
	})
	return _c
}

func (_c *MockCommenter_Comment_Call) Return(_a0 error) *MockCommenter_Comment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommenter_Comment_Call) RunAndReturn(run func(context.Context, PullRequestRepo, int, changeInfo) error) *MockCommenter_Comment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommenter creates a new instance of MockCommenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommenter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommenter {
	mock := &MockCommenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
