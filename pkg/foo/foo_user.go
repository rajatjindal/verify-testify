package foo

import (
	"fmt"
)

type foouser struct {
	foo Foo
}

func (f *foouser) RandomFoo(index int, inp string) {
	if index%2 == 0 {
		fmt.Println("calling Function one before two")
		f.foo.FunctionOne(inp)
		f.foo.FunctionTwo(inp)
		return
	}

	fmt.Println("calling Function two before one")
	f.foo.FunctionTwo(inp)
	f.foo.FunctionOne(inp)
}
