package foo

type Foo interface {
	FunctionOne(inp string) string
	FunctionTwo(inp string) string
}

type fooImpl struct{}

func (f *fooImpl) FunctionOne(inp string) string {
	return "foo impl function one"
}

func (f *fooImpl) FunctionTwi(inp string) string {
	return "foo impl function two"
}
