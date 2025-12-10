package core

import (
	"errors"
	"shared/pkg/data"
	"shared/pkg/types"
)

var or = handler(func(ctx *functionCtx) {
	var first = ctx.Stack.Pop()
	var second = ctx.Stack.Pop()

	if first.Type == types.LoInt && second.Type == types.LoInt {
		v, err := data.NewValue(first.GetInt() | second.GetInt())
		if err != nil {
			ctx.Vm.error(err)
		}
		ctx.Stack.Push(v)
		return
	}

	ctx.Vm.error(errors.New("unable to or the specified type pattern"))
})
