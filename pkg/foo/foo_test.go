package foo

import "testing"

func TestFoo(t *testing.T) {
	testfoo := &mockFoo{}

	testfoo.On("FunctionOne", "rj").Return("from mock function one")
	testfoo.On("FunctionTwo", "rj").Return("from mock function two")

	user := foouser{
		foo: testfoo,
	}

	user.RandomFoo(1, "rj")
	testfoo.AssertCalled(t, "FunctionOne", "rj")
	testfoo.AssertCalled(t, "FunctionTwo", "rj")

	user.RandomFoo(1, "rj")
	testfoo.AssertCalled(t, "FunctionTwo", "rj")
	testfoo.AssertCalled(t, "FunctionOne", "rj")

	user.RandomFoo(2, "rj")
	testfoo.AssertCalled(t, "FunctionOne", "rj")
	testfoo.AssertCalled(t, "FunctionTwo", "rj")

	user.RandomFoo(2, "rj")
	testfoo.AssertCalled(t, "FunctionTwo", "rj")
	testfoo.AssertCalled(t, "FunctionOne", "rj")
}
