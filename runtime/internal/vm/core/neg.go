package core

import (
	"errors"
	"shared/pkg/data"
	"shared/pkg/types"
)

var neg = handler(func(ctx *functionCtx) {
	var num = ctx.Stack.Pop()

	if num.Type == types.LoInt {
		v, err := data.NewValue(-num.GetInt())
		if err != nil {
			ctx.Vm.error(err)
		}
		ctx.Stack.Push(v)
		return
	}

	ctx.Vm.error(errors.New("unable to negate the specified type"))
})
