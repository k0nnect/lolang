package core

import (
	"errors"
	"shared/pkg/data"
	"shared/pkg/types"
)

var not = handler(func(ctx *functionCtx) {
	var num = ctx.Stack.Pop()

	if num.Type == types.LoInt {
		v, err := data.NewValue(^num.GetInt())
		if err != nil {
			ctx.Vm.error(err)
		}
		ctx.Stack.Push(v)
		return
	} else if num.Type == types.LoBool {
		v, err := data.NewValue(!num.GetBool())
		if err != nil {
			ctx.Vm.error(err)
		}
		ctx.Stack.Push(v)
		return
	}

	ctx.Vm.error(errors.New("unable to not the specified type"))
})
