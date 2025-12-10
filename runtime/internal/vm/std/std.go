package std

import (
	"shared/pkg/data"
	"shared/pkg/function"
)

type stdFunc struct {
	Name      string
	Execute   func(values []data.Value) data.Value
	Arguments []function.Argument
}

var StdLib = map[int]stdFunc{
	0: printFunc,
}
