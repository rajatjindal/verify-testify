package foo

import (
	"github.com/stretchr/testify/mock"
)

type mockFoo struct {
	mock.Mock
}

func (f *mockFoo) FunctionOne(inp string) string {
	return f.Called(inp).String(0)
}

func (f *mockFoo) FunctionTwo(inp string) string {
	return f.Called(inp).String(0)
}
