package function

import (
	"shared/pkg/data"
	"shared/pkg/function"
)

type Ctx struct {
	Function  *function.Function
	Stack     data.Stack[data.Value]
	Locals    []data.Value
	Arguments []data.Value
}
